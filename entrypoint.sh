#!/bin/sh

/app/fill-data -typesense $TYPESENSE_URL -token $TYPESENSE_APIKEY
/app/report-search \
  -typesense $TYPESENSE_URL \
  -token $TYPESENSE_APIKEY \
  -listen :80
