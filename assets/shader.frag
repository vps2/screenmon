#version 330 core

uniform float x1;
uniform float y1;
uniform float x2;
uniform float y2;
uniform float border;
uniform float height;

out vec4 color;

void main() {
    if(((gl_FragCoord.x > x1+border && gl_FragCoord.x < x2-border) || 
        (gl_FragCoord.x < x1-border && gl_FragCoord.x > x2+border)) &&
       ((gl_FragCoord.y > height-y1+border && gl_FragCoord.y < height-y2-border) || 
        (gl_FragCoord.y < height-y1-border && gl_FragCoord.y > height-y2+border)))
    {
        color = vec4(1.0f, 1.0f, 1.0f, 1.0f);
    } else {
        color = vec4(0.0f, 0.0f, 0.0f, 1.0f);
    }
}