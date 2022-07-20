#pragma once

#ifdef __cplusplus
extern "C" {
#endif


// igpPanic calls Go's panic() with the given string.
__attribute((__noreturn__)) void igpPanic(const char *msg);


#ifdef __cplusplus
}
#endif
