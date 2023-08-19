#!/bin/sh

/usr/bin/fill-data -typesense $TYPESENSE_URL -token $TYPESENSE_APIKEY
/usr/bin/report-search \
  -typesense $TYPESENSE_URL \
  -token $TYPESENSE_APIKEY \
  -listen :80
