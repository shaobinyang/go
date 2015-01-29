/* Created by cgo - DO NOT EDIT. */
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int), void *, int);

extern void _cgoexp_0d4ddd74268d_messageCallback1(void *, int);

void messageCallback1(int p0, void* p1, void* p2)
{
	struct {
		int p0;
		char __pad0[4];
		void* p1;
		void* p2;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	crosscall2(_cgoexp_0d4ddd74268d_messageCallback1, &a, 24);
}
