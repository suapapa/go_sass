#include "libsass_wrapper.h"

sass_context_t* _sass_new_context(void)
{
    return (sass_context_t*) sass_new_context();
}

sass_file_context_t* _sass_new_file_context(void)
{
    return (sass_file_context_t*) sass_new_file_context();
}

#if 0
sass_folder_context_t* _sass_new_folder_context(void)
{
    return (sass_folder_context_t*) sass_new_folder_context();
}
#endif

void _sass_free_context(sass_context_t* ctx)
{
    sass_free_context(ctx);
}

void _sass_free_file_context(sass_file_context_t* ctx)
{
    sass_free_file_context(ctx);
}

#if 0
void _sass_free_folder_context(sass_folder_context_t* ctx)
{
    sass_free_folder_context(ctx);
}
#endif

int _sass_compile(sass_context_t* ctx)
{
    return sass_compile(ctx);
}

int _sass_compile_file(sass_file_context_t* ctx)
{
    return sass_compile_file(ctx);
}

#if 0
int _sass_compile_folder(sass_folder_context_t* ctx)
{
    return sass_compile_folder(ctx);
}
#endif
