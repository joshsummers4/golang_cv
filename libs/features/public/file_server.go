package public

import (
	"crypto/md5"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed *.js cv/**.svg cv/**.ico
var embedded embed.FS

var etags map[string]string

func init() {
	etags = make(map[string]string)

	AddETags(embedded)
}

func AddETags(embedded embed.FS) {
	err := fs.WalkDir(embedded, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			content, err := embedded.ReadFile(path)
			if err != nil {
				log.Printf("Error reading embedded file %s: %v", path, err)
				return err
			}

			// Calculate the MD5 hash as the ETag
			hash := md5.Sum(content)
			etag := fmt.Sprintf(`"%x"`, hash)

			// Store the ETag for this file path
			etags["/"+path] = etag
		}
		return nil
	})

	if err != nil {
		log.Fatalf("\033[1;31m⛔  Error initializing ETags:\033[0m %v", err)
	}
}

func cachePolicyHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the requested file path
		filename := r.URL.Path

		// Get the precomputed ETag for the file
		etag, ok := etags[filename]
		if !ok {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// Set ETag header
		w.Header().Set("ETag", etag)

		// Handle If-None-Match header to check if the ETag matches
		if match := r.Header.Get("If-None-Match"); match == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		// Set other caching headers (e.g., Cache-Control for browser caching)
		w.Header().Set("Cache-Control", "no-cache") // Cache for 1 year

		// Call the original handler to serve the file
		h.ServeHTTP(w, r)
	})
}

func CreateFileServer(fs embed.FS) http.Handler {
	sharedHandler := http.FileServer(http.FS(embedded))
	localHandler := http.FileServer(http.FS(fs))

	combinedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to serve from shared filesystem first
		path := strings.TrimPrefix(r.URL.Path, "/")
		sharedFile, err := embedded.Open(path)
		if err == nil {
			sharedFile.Close()
			sharedHandler.ServeHTTP(w, r)
			return
		}

		localFile, err := fs.Open(path)
		if err == nil {
			localFile.Close()
			localHandler.ServeHTTP(w, r)
			return
		}

		// If not found in either filesystem, return 404
		http.Error(w, "File not found", http.StatusNotFound)
	})

	return cachePolicyHandler(combinedHandler)
}
