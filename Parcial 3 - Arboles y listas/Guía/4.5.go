// ejercicio de Mario
Dada una lista simplemente encadenada que contiene datos de todas las provincias de la República Argentina: nombre, capital, cantidad total de habitantes y cantidad de analfabetos, y está ordenada en forma decreciente por número de habitantes analfabetos, generar otras tres listas que contengan el nombre, la capital y el porcentaje de analfabetos de las Provincias que respondan a las siguientes restricciones.

L1: <= 10 % analfabetos
L2: 16 a 25 % analfabetos
L3: => 26 % analfabetos

accion 4.5 (prim:puntero a nodo) es 
	ambiente
		nodo=regsitro 	
			nombre:an(30)
			capital:an(30)
			can_total:N(4)
			can_analfa:N(4)
			prox:puntero a nodo 
		finregistro
		nodo_s=regsitro 	
			nombre:an(30)
			capital:an(30)
			porcentajee:N(4)
			prox1:puntero a nodo 
		finregistro
		nodo1,prim1,prim2,prim3:puntero a nodo_s
		nod,prim:putero a nodo
		l2:(16..25)

		Procedimiento carga() es//como los 3 tiene el mismo registro solo hace falta hacer 1 vez y usar un solo nodo1
			nuevo(nodo1)
			*nodo1.nombre:=*nod.nombre
			*nodo1.capital:=*nod.capital
			*nodo1.procentajee:=procentaje
		finprocedimiento
		proceso  
		procentaje=(nod.can_analf / nod.can_total) * 100
		nod:=prim 
		prim1:=nil
		prim2:=nil
		prim3:=nil
		si prim=nil entonces//verifico si la lista no esta vacia
			esc"la lista esta vacia"
		sino 
			mientras nod<>nil hacer //recorro 

				Si procentaje <=10 entonces
					carga()//aca recien cargo mi nodito 
					*nodo1.prox1:=prim1//asi hago la conexion con el nodo
					prim1:=nodo1
				sino
					si procentaje >=16 o procentaje <=25 entonces
						carga()
						*nodo1.prox1:=prim2//asi hago la conexion con el nodo
						prim2:=nodo1
					sino 
						si procentaje >=26 entonces 
							carga()
							*nodo1.prox1:=prim3//asi hago la conexion con el nodo
							prim3:=nodo1
						FinSi
					FinSi
				finsi
				nod:=*nod.prox//avzo el nodo 
			finmientras
		finsi
finaccion

	