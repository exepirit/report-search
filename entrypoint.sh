#!/bin/sh

/app/report-search \
  -typesense $TYPESENSE_URL \
  -token $TYPESENSE_APIKEY \
  -listen :80
