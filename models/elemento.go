package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Elemento struct {
	Id                 int       `orm:"column(id_elemento);pk"`
	FechaRegistro      time.Time `orm:"column(fecha_registro);type(date)"`
	Nivel              int       `orm:"column(nivel)"`
	TipoBien           int       `orm:"column(tipo_bien)"`
	Descripcion        string    `orm:"column(descripcion)"`
	Cantidad           float64   `orm:"column(cantidad)"`
	Unidad             string    `orm:"column(unidad);null"`
	Valor              float64   `orm:"column(valor);null"`
	Iva                float64   `orm:"column(iva);null"`
	Ajuste             float64   `orm:"column(ajuste);null"`
	Bodega             int       `orm:"column(bodega);null"`
	SubtotalSinIva     float64   `orm:"column(subtotal_sin_iva);null"`
	TotalIva           float64   `orm:"column(total_iva);null"`
	TotalIvaCon        float64   `orm:"column(total_iva_con);null"`
	TipoPoliza         int       `orm:"column(tipo_poliza);null"`
	FechaInicioPol     time.Time `orm:"column(fecha_inicio_pol);type(date);null"`
	FechaFinalPol      time.Time `orm:"column(fecha_final_pol);type(date);null"`
	Marca              string    `orm:"column(marca);null"`
	Serie              string    `orm:"column(serie);null"`
	IdEntrada          int       `orm:"column(id_entrada)"`
	Estado             bool      `orm:"column(estado)"`
	CantidadPorAsignar float64   `orm:"column(cantidad_por_asignar);null"`
}

func (t *Elemento) TableName() string {
	return "elemento"
}

func init() {
	orm.RegisterModel(new(Elemento))
}

// AddElemento insert a new Elemento into database and returns
// last inserted Id on success.
func AddElemento(m *Elemento) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetElementoById retrieves Elemento by Id. Returns error if
// Id doesn't exist
func GetElementoById(id int) (v *Elemento, err error) {
	o := orm.NewOrm()
	v = &Elemento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllElemento retrieves all Elemento matches certain condition. Returns empty list if
// no records exist
func GetAllElemento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Elemento))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Elemento
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateElemento updates Elemento by Id and returns error if
// the record to be updated doesn't exist
func UpdateElementoById(m *Elemento) (err error) {
	o := orm.NewOrm()
	v := Elemento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteElemento deletes Elemento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteElemento(id int) (err error) {
	o := orm.NewOrm()
	v := Elemento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Elemento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
