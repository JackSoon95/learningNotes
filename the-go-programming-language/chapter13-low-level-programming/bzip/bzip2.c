/* This file is gopl.io/ch13/bzip/bzip2.c, */
/* a simple wrapper for libbzip2 suitable for cgo. */
#include <bzlib.h>

int bz2compress(bz_stream *s, int action, 
                char *in, unsigned *inlen, char *out, unsigned *outlen) {
  s->next_in = in;
  s->avail_in = *inlen;
  s->next_out = out;
  s->avail_out = *outlenm;
  int r = BZ2_bzCompress(s, action) ;
  *inlen -= s->avail_in;
  *outlen -= s->avail_out;
  return r
}
