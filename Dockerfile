FROM scratch
COPY server ./server
COPY static ./static
CMD ["./server"]