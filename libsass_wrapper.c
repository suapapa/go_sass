#include "libsass_wrapper.h"

sass_context_t* _sass_new_context(void)
{
	return (sass_context_t*)sass_new_context();
}

void _sass_free_context(sass_context_t* ctx)
{
	sass_free_context(ctx);
}

int _sass_compile(sass_context_t* ctx)
{
	return sass_compile(ctx);
}

