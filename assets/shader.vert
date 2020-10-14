#version 330 core  
 
layout (location = 0) in vec3 position;

//ширина и высота окна
uniform float width;
uniform float height;

void main() {
    //нормализуем координаты окна в локальные opengl 
    gl_Position = vec4(position.x * 2.0 / width - 1.0, 
                       position.y * -2.0 / height + 1.0, 
                       position.z, 
                       1.0);
}