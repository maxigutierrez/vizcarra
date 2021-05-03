

select b.id,mediodepago,max(c.id) from caja c inner join 
recibos_facturas a on c.comprobante_id = a.recibos_id inner join 
comprobantes  b on a.comprobantes_id = b.id
where b.total = abs(ingreso+egreso)
group by cajas_id,b.id,mediodepago
having sum(egreso)>sum(ingreso)





select c.*,(select franquicia from franquicias where id = b.franquicia_cobro) from caja c inner join 
recibos_facturas a on c.comprobante_id = a.recibos_id inner join 
comprobantes  b on a.comprobantes_id = b.id
where  b.id=468598


    INSERT INTO [dbo].[caja]
            ([cajas_id]
            ,[fecha]
            ,[descripcion]
            ,[mediodepago]
            ,[comprobante]
            ,[tipo]
            ,[comprobante_id]
            ,[ingreso]
            ,[egreso]
            ,[cierres_id])
    SELECT 
        [cajas_id]
        ,getdate()
        ,'Ajuste Cobro Anulado'
        ,[mediodepago]
        ,[comprobante]
        ,[tipo]
        ,comprobante_id
        ,0
        ,egreso*-1
        ,null
    FROM [dbo].[caja]
    where id = 95871


