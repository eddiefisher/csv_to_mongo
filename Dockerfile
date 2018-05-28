FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY /bin/csv_to_mongo /
CMD ["/csv_to_mongo"]