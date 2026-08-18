package protocol

import "fmt"

func FrameMessage(message string, username string) [][]byte {
	lastChunkSize, numOfChunks := chooseChunks(len(message))
	// Makes/allocates the array for the byte arrays
	chunks := make([][]byte, numOfChunks)
	// Sets chunks to byte arrays with this format: chunk_iduusernameddata
	for chunk := 0; chunk < numOfChunks-1; chunk++ {
		chunks[chunk] = []byte(fmt.Sprintf("%vu%sd%s", chunk, username, message[1000*chunk:1000*chunk+1]))
	}
	chunks[numOfChunks-1] = []byte(fmt.Sprintf("%vu%sd%s", numOfChunks-1, username, message[numOfChunks*1000:numOfChunks*1000+lastChunkSize]))
	return chunks
}

// Calculates the needed amount of chunks and the last one(reminder), currently every chunk is 1 kb
func chooseChunks(len int) (int, int) {
	const chunkSize = 1000
	// Less than a kb not worth chunking
	if len < chunkSize {
		return len, 1
	}
	numOfChunks := len / chunkSize
	lastChunk := len % chunkSize

	// If there is no reminder no need for last chunk
	if lastChunk == 0 {
		return chunkSize, numOfChunks
	}

	// + 1 for the chunk reminder
	return lastChunk, numOfChunks + 1
}
