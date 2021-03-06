
#include "v8_c_private.h"

#include <cstdlib>
#include <cstring>

extern "C" {
  Version version = {V8_MAJOR_VERSION, V8_MINOR_VERSION, V8_BUILD_NUMBER, V8_PATCH_LEVEL};

  void v8_Initialize() {
    const char* flags = "--expose_gc";
    v8::V8::SetFlagsFromString(flags, strlen(flags));

    platform = v8::platform::CreateDefaultPlatform();
    v8::V8::InitializePlatform(platform);
    v8::V8::Initialize();
    return;
  }
}
