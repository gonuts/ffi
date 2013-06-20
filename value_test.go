package ffi_test

import (
	"reflect"
	"testing"

	"github.com/gonuts/ffi"
)

func TestGetSetBuiltinValue(t *testing.T) {

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"int", ffi.C_int, int64(val)},
			{"int8", ffi.C_int8, int64(val)},
			{"int16", ffi.C_int16, int64(val)},
			{"int32", ffi.C_int32, int64(val)},
			{"int64", ffi.C_int64, int64(val)},
		} {
			cval := ffi.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Int(), cval.Int())
			cval.SetInt(val)
			eq(t, tt.val, cval.Int())
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"unsigned int", ffi.C_uint, uint64(val)},
			{"uint8", ffi.C_uint8, uint64(val)},
			{"uint16", ffi.C_uint16, uint64(val)},
			{"uint32", ffi.C_uint32, uint64(val)},
			{"uint64", ffi.C_uint64, uint64(val)},
		} {
			cval := ffi.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Uint(), cval.Uint())
			cval.SetUint(val)
			eq(t, tt.val, cval.Uint())
		}
	}

	{
		const val = -66.0
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"float", ffi.C_float, float64(val)},
			{"double", ffi.C_double, float64(val)},
			//FIXME: Go has no equivalent for long double...
			//{"long double", ffi.C_longdouble, float128(val)},
		} {
			cval := ffi.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Float(), cval.Float())
			cval.SetFloat(val)
			eq(t, tt.val, cval.Float())
		}
	}

	{
		const val = -66
		cval := ffi.New(ffi.C_int64)
		cptr := cval.Addr()
		cval.SetInt(val)
		eq(t, int64(val), cval.Int())
		eq(t, int64(val), cptr.Elem().Int())
		cval.SetInt(0)
		eq(t, int64(0), cptr.Elem().Int())
		cptr.Elem().SetInt(val)
		eq(t, int64(val), cval.Int())
		eq(t, int64(val), cptr.Elem().Int())

	}
}

func TestGetSetArrayValue(t *testing.T) {

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			len int
			t   ffi.Type
			val interface{}
		}{
			{"uint8[10]", 10, ffi.C_uint8, [10]uint8{}},
			{"uint16[10]", 10, ffi.C_uint16, [10]uint16{}},
			{"uint32[10]", 10, ffi.C_uint32, [10]uint32{}},
			{"uint64[10]", 10, ffi.C_uint64, [10]uint64{}},
		} {
			ctyp, err := ffi.NewArrayType(tt.len, tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			len int
			t   ffi.Type
			val interface{}
		}{
			{"int8[10]", 10, ffi.C_int8, [10]int8{}},
			{"int16[10]", 10, ffi.C_int16, [10]int16{}},
			{"int32[10]", 10, ffi.C_int32, [10]int32{}},
			{"int64[10]", 10, ffi.C_int64, [10]int64{}},
		} {
			ctyp, err := ffi.NewArrayType(tt.len, tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			len int
			t   ffi.Type
			val interface{}
		}{
			{"float[10]", 10, ffi.C_float, [10]float32{}},
			{"double[10]", 10, ffi.C_double, [10]float64{}},
			// FIXME: go has no long double equivalent
			//{"long double[10]", 10, ffi.C_longdouble, [10]float128{}},
		} {
			ctyp, err := ffi.NewArrayType(tt.len, tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}

}

func TestGetSetStructValue(t *testing.T) {

	const val = 42
	arr10, err := ffi.NewArrayType(10, ffi.C_int32)
	if err != nil {
		t.Errorf(err.Error())
	}

	ctyp, err := ffi.NewStructType(
		"struct_ssv",
		[]ffi.Field{
			{"F1", ffi.C_uint16},
			{"F2", arr10},
			{"F3", ffi.C_int32},
			{"F4", ffi.C_uint16},
		})
	eq(t, "struct_ssv", ctyp.Name())
	eq(t, ffi.Struct, ctyp.Kind())
	eq(t, 4, ctyp.NumField())

	cval := ffi.New(ctyp)
	eq(t, ctyp.Kind(), cval.Kind())
	eq(t, ctyp.NumField(), cval.NumField())
	eq(t, uint64(0), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(0), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(0), cval.Field(2).Int())
	eq(t, uint64(0), cval.Field(3).Uint())

	// set everything to 'val'
	cval.Field(0).SetUint(val)
	for i := 0; i < arr10.Len(); i++ {
		cval.Field(1).Index(i).SetInt(val)
	}
	cval.Field(2).SetInt(val)
	cval.Field(3).SetUint(val)

	// test values back
	eq(t, uint64(val), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(val), cval.Field(2).Int())
	eq(t, uint64(val), cval.Field(3).Uint())

	// test values back - by field name
	eq(t, uint64(val), cval.FieldByName("F1").Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.FieldByName("F2").Index(i).Int())
	}
	eq(t, int64(val), cval.FieldByName("F3").Int())
	eq(t, uint64(val), cval.FieldByName("F4").Uint())
}

func TestGetSetStructWithSliceValue(t *testing.T) {

	const val = 42
	arr10, err := ffi.NewArrayType(10, ffi.C_int32)
	if err != nil {
		t.Errorf(err.Error())
	}
	slityp, err := ffi.NewSliceType(ffi.C_int32)
	if err != nil {
		t.Errorf(err.Error())
	}

	ctyp, err := ffi.NewStructType(
		"struct_sswsv",
		[]ffi.Field{
			{"F1", ffi.C_uint16},
			{"F2", arr10},
			{"F3", ffi.C_int32},
			{"F4", ffi.C_uint16},
			{"F5", slityp},
		})
	eq(t, "struct_sswsv", ctyp.Name())
	eq(t, ffi.Struct, ctyp.Kind())
	eq(t, 5, ctyp.NumField())

	cval := ffi.New(ctyp)
	eq(t, ctyp.Kind(), cval.Kind())
	eq(t, ctyp.NumField(), cval.NumField())
	eq(t, uint64(0), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(0), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(0), cval.Field(2).Int())
	eq(t, uint64(0), cval.Field(3).Uint())
	eq(t, int(0), cval.Field(4).Len())
	eq(t, int(0), cval.Field(4).Len())

	goval := struct {
		F1 uint16
		F2 [10]int32
		F3 int32
		F4 uint16
		F5 []int32
	}{
		F1: val,
		F2: [10]int32{val, val, val, val, val,
			val, val, val, val, val},
		F3: val,
		F4: val,
		F5: make([]int32, 2, 3),
	}
	goval.F5[0] = val
	goval.F5[1] = val

	cval.SetValue(reflect.ValueOf(goval))

	eq(t, uint64(val), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(val), cval.Field(2).Int())
	eq(t, uint64(val), cval.Field(3).Uint())
	eq(t, int(2), cval.Field(4).Len())
	// FIXME: should we get the 'cap' from go ?
	eq(t, int( /*3*/ 2), cval.Field(4).Cap())
	eq(t, int64(val), cval.Field(4).Index(0).Int())
	eq(t, int64(val), cval.Field(4).Index(1).Int())
}

func TestGetSetSliceValue(t *testing.T) {

	const sz = 10
	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"uint8[]", ffi.C_uint8, make([]uint8, sz)},
			{"uint16[]", ffi.C_uint16, make([]uint16, sz)},
			{"uint32[]", ffi.C_uint32, make([]uint32, sz)},
			{"uint64[]", ffi.C_uint64, make([]uint64, sz)},
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"int8[]", ffi.C_int8, make([]int8, sz)},
			{"int16[]", ffi.C_int16, make([]int16, sz)},
			{"int32[]", ffi.C_int32, make([]int32, sz)},
			{"int64[]", ffi.C_int64, make([]int64, sz)},
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"float[]", ffi.C_float, make([]float32, sz)},
			{"double[]", ffi.C_double, make([]float64, sz)},
			// FIXME: go has no long double equivalent
			//{"long double[]", ffi.C_longdouble, make([]float128, sz)}
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}

	// now test if slices can automatically grow...
	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"uint8[]", ffi.C_uint8, make([]uint8, sz)},
			{"uint16[]", ffi.C_uint16, make([]uint16, sz)},
			{"uint32[]", ffi.C_uint32, make([]uint32, sz)},
			{"uint64[]", ffi.C_uint64, make([]uint64, sz)},
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"int8[]", ffi.C_int8, make([]int8, sz)},
			{"int16[]", ffi.C_int16, make([]int16, sz)},
			{"int32[]", ffi.C_int32, make([]int32, sz)},
			{"int64[]", ffi.C_int64, make([]int64, sz)},
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			t   ffi.Type
			val interface{}
		}{
			{"float[]", ffi.C_float, make([]float32, sz)},
			{"double[]", ffi.C_double, make([]float64, sz)},
			// FIXME: go has no long double equivalent
			//{"long double[]", ffi.C_longdouble, make([]float128, sz)}
		} {
			ctyp, err := ffi.NewSliceType(tt.t)
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := ffi.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}
}

func TestValueOf(t *testing.T) {
	{
		const val = 42
		for _, v := range []interface{}{
			int(val),
			int8(val),
			int16(val),
			int32(val),
			int64(val),
		} {
			eq(t, int64(val), ffi.ValueOf(v).Int())
		}
	}

	{
		const val = 42
		for _, v := range []interface{}{
			uint(val),
			uint8(val),
			uint16(val),
			uint32(val),
			uint64(val),
		} {
			eq(t, uint64(val), ffi.ValueOf(v).Uint())
		}
	}
	{
		const val = 42.0
		for _, v := range []interface{}{
			float32(val),
			float64(val),
		} {
			eq(t, float64(val), ffi.ValueOf(v).Float())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewStructType(
			"struct_ints",
			[]ffi.Field{
				{"F1", ffi.C_int8},
				{"F2", ffi.C_int16},
				{"F3", ffi.C_int32},
				{"F4", ffi.C_int64},
			})
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		for i := 0; i < ctyp.NumField(); i++ {
			cval.Field(i).SetInt(int64(val))
			eq(t, int64(val), cval.Field(i).Int())
		}
		gval := struct {
			F1 int8
			F2 int16
			F3 int32
			F4 int64
		}{val + 1, val + 1, val + 1, val + 1}
		rval := reflect.ValueOf(gval)
		eq(t, rval.NumField(), cval.NumField())
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int()-1, cval.Field(i).Int())
		}
		cval = ffi.ValueOf(gval)
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int(), cval.Field(i).Int())
		}
	}
}

func TestEncoderDecoder(t *testing.T) {
	arr_10, _ := ffi.NewArrayType(10, ffi.C_int32)
	sli_10, _ := ffi.NewSliceType(ffi.C_int32)

	const sz = 10
	{
		const val = 42
		for _, v := range []interface{}{
			int(val),
			int8(val),
			int16(val),
			int32(val),
			int64(val),
		} {
			ct := ffi.TypeOf(v)
			cv := ffi.New(ct)
			enc := ffi.NewEncoder(cv)
			err := enc.Encode(v)
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, int64(val), cv.Int())

			// now decode back
			vv := reflect.New(reflect.TypeOf(v))
			dec := ffi.NewDecoder(cv)
			err = dec.Decode(vv.Interface())
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, vv.Elem().Int(), cv.Int())
		}
	}

	{
		const val = 42
		for _, v := range []interface{}{
			uint(val),
			uint8(val),
			uint16(val),
			uint32(val),
			uint64(val),
		} {
			ct := ffi.TypeOf(v)
			cv := ffi.New(ct)
			enc := ffi.NewEncoder(cv)
			err := enc.Encode(v)
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, uint64(val), cv.Uint())

			// now decode back
			vv := reflect.New(reflect.TypeOf(v))
			dec := ffi.NewDecoder(cv)
			err = dec.Decode(vv.Interface())
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, vv.Elem().Uint(), cv.Uint())
		}
	}
	{
		const val = 42.0
		for _, v := range []interface{}{
			float32(val),
			float64(val),
		} {
			ct := ffi.TypeOf(v)
			cv := ffi.New(ct)
			enc := ffi.NewEncoder(cv)
			err := enc.Encode(v)
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, float64(val), cv.Float())

			// now decode back
			vv := reflect.New(reflect.TypeOf(v))
			dec := ffi.NewDecoder(cv)
			err = dec.Decode(vv.Interface())
			if err != nil {
				t.Errorf(err.Error())
			}
			eq(t, vv.Elem().Float(), cv.Float())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewStructType(
			"struct_ints",
			[]ffi.Field{
				{"F1", ffi.C_int8},
				{"F2", ffi.C_int16},
				{"F3", ffi.C_int32},
				{"F4", ffi.C_int64},
			})
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		gval := struct {
			F1 int8
			F2 int16
			F3 int32
			F4 int64
		}{val + 1, val + 1, val + 1, val + 1}
		err = ffi.Associate(ctyp, reflect.TypeOf(gval))
		if err != nil {
			t.Errorf(err.Error())
		}

		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}

		rval := reflect.ValueOf(gval)
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int(), cval.Field(i).Int())
		}

		// now decode back
		vv := reflect.New(rval.Type())
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		rval = vv.Elem()
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int(), cval.Field(i).Int())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewStructType(
			"struct_ints_arr10",
			[]ffi.Field{
				{"F1", ffi.C_int8},
				{"F2", ffi.C_int16},
				{"A1", arr_10},
				{"F3", ffi.C_int32},
				{"F4", ffi.C_int64},
			})
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		gval := struct {
			F1 int8
			F2 int16
			A1 [sz]int32
			F3 int32
			F4 int64
		}{
			val + 1, val + 1,
			[sz]int32{
				val, val, val, val, val,
				val, val, val, val, val,
			},
			val + 1, val + 1,
		}
		err = ffi.Associate(ctyp, reflect.TypeOf(gval))
		if err != nil {
			t.Errorf(err.Error())
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}

		rval := reflect.ValueOf(gval)
		eq(t, rval.Field(0).Int(), cval.Field(0).Int())
		eq(t, rval.Field(1).Int(), cval.Field(1).Int())
		eq(t, rval.Field(3).Int(), cval.Field(3).Int())
		eq(t, rval.Field(4).Int(), cval.Field(4).Int())
		rfield := cval.Field(2)
		cfield := cval.Field(2)
		eq(t, rfield.Len(), cfield.Len())
		for i := 0; i < cfield.Len(); i++ {
			eq(t, rfield.Index(i).Int(), cfield.Index(i).Int())
		}

		// now decode back
		vv := reflect.New(rval.Type())
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		rval = vv.Elem()
		eq(t, rval.Field(0).Int(), cval.Field(0).Int())
		eq(t, rval.Field(1).Int(), cval.Field(1).Int())
		eq(t, rval.Field(3).Int(), cval.Field(3).Int())
		eq(t, rval.Field(4).Int(), cval.Field(4).Int())
		rfield = cval.Field(2)
		cfield = cval.Field(2)
		eq(t, rfield.Len(), cfield.Len())
		for i := 0; i < cfield.Len(); i++ {
			eq(t, rfield.Index(i).Int(), cfield.Index(i).Int())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewStructType(
			"struct_ints_sli10",
			[]ffi.Field{
				{"F1", ffi.C_int8},
				{"F2", ffi.C_int16},
				{"S1", sli_10},
				{"F3", ffi.C_int32},
				{"F4", ffi.C_int64},
			})
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		gval := struct {
			F1 int8
			F2 int16
			S1 []int32
			F3 int32
			F4 int64
		}{
			val + 1, val + 1,
			[]int32{
				val, val, val, val, val,
				val, val, val, val, val,
			},
			val + 1, val + 1,
		}
		err = ffi.Associate(ctyp, reflect.TypeOf(gval))
		if err != nil {
			t.Errorf(err.Error())
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}

		rval := reflect.ValueOf(gval)
		eq(t, rval.Field(0).Int(), cval.Field(0).Int())
		eq(t, rval.Field(1).Int(), cval.Field(1).Int())
		eq(t, rval.Field(3).Int(), cval.Field(3).Int())
		eq(t, rval.Field(4).Int(), cval.Field(4).Int())
		rfield := cval.Field(2)
		cfield := cval.Field(2)
		eq(t, rfield.Len(), cfield.Len())
		for i := 0; i < cfield.Len(); i++ {
			eq(t, rfield.Index(i).Int(), cfield.Index(i).Int())
		}

		// now decode back
		vv := reflect.New(rval.Type())
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		rval = vv.Elem()
		eq(t, rval.Field(0).Int(), cval.Field(0).Int())
		eq(t, rval.Field(1).Int(), cval.Field(1).Int())
		eq(t, rval.Field(3).Int(), cval.Field(3).Int())
		eq(t, rval.Field(4).Int(), cval.Field(4).Int())
		rfield = cval.Field(2)
		cfield = cval.Field(2)
		eq(t, rfield.Len(), cfield.Len())
		for i := 0; i < cfield.Len(); i++ {
			eq(t, rfield.Index(i).Int(), cfield.Index(i).Int())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewArrayType(sz, ffi.C_int32)
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		gval := [sz]int32{
			val, val, val, val, val,
			val, val, val, val, val,
		}
		err = ffi.Associate(ctyp, reflect.TypeOf(gval))
		if err != nil {
			t.Errorf(err.Error())
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Type().Len(); i++ {
			eq(t, int64(val), cval.Index(i).Int())
		}

		// now decode back
		vv := reflect.New(reflect.TypeOf(gval))
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Type().Len(); i++ {
			eq(t, vv.Elem().Index(i).Int(), cval.Index(i).Int())
		}

	}

	{
		const val = 42
		ctyp, err := ffi.NewArrayType(sz, ffi.C_float)
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.New(ctyp)
		gval := [sz]float32{
			val, val, val, val, val,
			val, val, val, val, val,
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Type().Len(); i++ {
			eq(t, float64(val), cval.Index(i).Float())
		}
		// now decode back
		vv := reflect.New(reflect.TypeOf(gval))
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Type().Len(); i++ {
			eq(t, vv.Elem().Index(i).Float(), cval.Index(i).Float())
		}

	}

	{
		const val = 42
		ctyp, err := ffi.NewSliceType(ffi.C_int32)
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.MakeSlice(ctyp, sz, sz)
		gval := []int32{
			val, val, val, val, val,
			val, val, val, val, val,
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Len(); i++ {
			eq(t, int64(val), cval.Index(i).Int())
		}
		// now decode back
		vv := reflect.New(reflect.TypeOf(gval))
		dec := ffi.NewDecoder(cval)
		err = dec.Decode(vv.Interface())
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Len(); i++ {
			eq(t, vv.Elem().Index(i).Int(), cval.Index(i).Int())
		}
	}
	{
		const val = 42
		ctyp, err := ffi.NewSliceType(ffi.C_float)
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := ffi.MakeSlice(ctyp, sz, sz)
		gval := []float32{
			val, val, val, val, val,
			val, val, val, val, val,
		}
		enc := ffi.NewEncoder(cval)
		err = enc.Encode(gval)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i := 0; i < cval.Len(); i++ {
			eq(t, float64(val), cval.Index(i).Float())
		}
		{
			// now decode back
			vv := reflect.New(reflect.TypeOf(gval))
			dec := ffi.NewDecoder(cval)
			err = dec.Decode(vv.Interface())
			if err != nil {
				t.Errorf(err.Error())
			}
			for i := 0; i < cval.Len(); i++ {
				eq(t, vv.Elem().Index(i).Float(), cval.Index(i).Float())
			}
		}
	}
}

func TestAllocValueOf(t *testing.T) {
	const nmax = 10000
	type Event struct {
		F   float64
		Arr [2]float64
		Sli []float64
	}
	type branch struct {
		g      reflect.Value
		c      ffi.Value
		update func()
	}
	var evt Event
	evt.Sli = make([]float64, 0)
	set_branch := func(objaddr interface{}) *branch {
		ptr := reflect.ValueOf(objaddr)
		val := reflect.Indirect(ptr)
		cval := ffi.ValueOf(val.Interface())
		var br *branch
		br = &branch{
			g: val,
			c: cval,
			update: func() {
				br.c.SetValue(br.g)
			},
		}
		return br
	}
	br := set_branch(&evt)

	for i := 0; i < nmax; i++ {
		evt.F = float64(i + 1)
		evt.Arr[0] = -evt.F
		evt.Arr[1] = -2 * evt.F
		evt.Sli = evt.Sli[:0]
		evt.Sli = append(evt.Sli, -evt.F)
		evt.Sli = append(evt.Sli, -2*evt.F)

		br.update()
		eq(t, evt.F, br.c.Field(0).Float())
		eq(t, evt.Arr[0], br.c.Field(1).Index(0).Float())
		eq(t, evt.Arr[1], br.c.Field(1).Index(1).Float())
		eq(t, evt.Sli[0], br.c.Field(2).Index(0).Float())
		eq(t, evt.Sli[1], br.c.Field(2).Index(1).Float())
	}

	for i := 0; i < nmax; i++ {
		evt.F = float64(i + 1)
		evt.Arr[0] = -evt.F
		evt.Arr[1] = -2 * evt.F
		evt.Sli[0] = -evt.F
		evt.Sli[1] = -2 * evt.F

		br.update()
		eq(t, evt.F, br.c.Field(0).Float())
		eq(t, evt.Arr[0], br.c.Field(1).Index(0).Float())
		eq(t, evt.Arr[1], br.c.Field(1).Index(1).Float())
		eq(t, evt.Sli[0], br.c.Field(2).Index(0).Float())
		eq(t, evt.Sli[1], br.c.Field(2).Index(1).Float())
	}

}

// EOF
