/*
Copyright © 2012 Serge Zirukin

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

#ifndef _GOCHIPMUNK_SPACE_H
#define _GOCHIPMUNK_SPACE_H

inline cpBool space_add_poststep(cpSpace *space, cpDataPointer key, cpDataPointer data);
inline void space_add_collision_handler(cpSpace *space, cpCollisionType a, cpCollisionType b);
inline void space_set_default_collision_handler(cpSpace *space);
inline void space_bb_query(cpSpace *space, cpBB bb, cpLayers layers, cpGroup group, void *f);
inline void space_each_body(cpSpace *space, void *f);
inline void space_each_constraint(cpSpace *space, void *f);
inline void space_each_shape(cpSpace *space, void *f);
inline void space_nearest_point_query(cpSpace *space, cpVect point, cpFloat maxDistance,
                                      cpLayers layers, cpGroup group, void *f);
inline void space_point_query(cpSpace *s, cpVect point, cpLayers layers, cpGroup group, void *p);
inline void space_segment_query(cpSpace *space,
                                cpVect   start,
                                cpVect   end,
                                cpLayers layers,
                                cpGroup  group,
                                void    *f);

#endif // !_GOCHIPMUNK_SPACE_H

// Local Variables:
// indent-tabs-mode: nil
// c-basic-offset: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
