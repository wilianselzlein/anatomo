SELECT distinct cod_item, descricao 
FROM itens_agrupados ia 
join procedimento_guia pg on
	pg.cd_procedimento = ia.cod_item 
join guia_cobranca gc on 
	gc.CD_GUIA = pg.CD_GUIA_COBRANCA 
where 
	gc.DS_GUIA_STATUS = 'GUIA_VALIDA' 
	and pg.dt_execucao >= '|start|' 
	and pg.dt_execucao < '|end|' 
	and upper(grupo) in ('SAD', 'HONORARIO MEDICO CIRURGICO')
	and gc.cd_guia in (
		select distinct gc_ia.cd_guia
		from guia_cobranca gc_ia
 		join procedimento_guia pg_ia on
			gc_ia.CD_GUIA = pg_ia.CD_GUIA_COBRANCA
		where gc_ia.DS_GUIA_STATUS = 'GUIA_VALIDA'
			and pg_ia.dt_execucao >= '|start|'
			and pg_ia.dt_execucao < '|end|'
			and pg_ia.cd_procedimento = %d )
order by cod_item
