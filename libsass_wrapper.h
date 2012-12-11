#ifndef __LIBSASS_WRAPPER_H__
#define __LIBSASS_WRAPPER_H__
#include "sass_interface.h"

typedef struct sass_options sass_options_t;

typedef struct sass_context sass_context_t;
sass_context_t* _sass_new_context(void);
void _sass_free_context(sass_context_t* ctx);
int _sass_compile(sass_context_t* ctx);

typedef struct sass_file_context sass_file_context_t;
sass_file_context_t* _sass_new_file_context(void);
void _sass_free_file_context(sass_file_context_t* ctx);
int _sass_compile_file(sass_file_context_t* ctx);

#if 0 // compile_folder is not implemented in libsass
typedef struct sass_folder_context sass_folder_context_t;
sass_folder_context_t* _sass_new_folder_context(void);
void _sass_free_folder_context(sass_folder_context_t* ctx);
int _sass_compile_folder(sass_folder_context_t* ctx);
#endif

#endif /* __LIBSASS_WRAPPER_H__ */
