#ifndef __LIBSASS_WRAPPER_H__
#define __LIBSASS_WRAPPER_H__
#include "sass_interface.h"

typedef struct sass_options sass_options_t;
typedef struct sass_context sass_context_t;

sass_context_t* _sass_new_context(void);
void _sass_free_context(sass_context_t* ctx);
int _sass_compile(sass_context_t* ctx);
#endif /* __LIBSASS_WRAPPER_H__ */
