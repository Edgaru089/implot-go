#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif


// implot.CreateContext() [Contexts.go]
igpContext igpCreateContext();

// implot.(*Context).Destroy() [Contexts.go]
void igpDestroyContext(igpContext ctx);

// implot.CurrentContext() [Contexts.go]
igpContext igpCurrentContext();

// implot.(*Context).SetCurrent() [Contexts.go]
void igpSetCurrentContext(igpContext ctx);

// implot.SetImGUIContext() [Contexts.go]
void igpSetImGUIContext(iggContext ctx);


#ifdef __cplusplus
}
#endif
