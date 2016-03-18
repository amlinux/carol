#include "glad.h"
#include "sdl.h"
#include "render.h"

int InitOpenGL() {
	if(!gladLoadGLLoader(SDL_GL_GetProcAddress)) {
		return -1;
	}
	glClearColor (0.45, 0.31, 0.59, 1.0);
	return 0;
}

int Render(SDL_Window* w) {
	glClear (GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
	SDL_GL_SwapWindow(w);
	return 0;
}