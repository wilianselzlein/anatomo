with 
	item_agrupado as (
		select distinct
			gc.cd_guia as ia_cd_guia, pg.cd_procedimento as ia_cd_procedimento, pg.dt_execucao as ia_dt_execucao
		from
			guia_cobranca gc
		join procedimento_guia pg on
			gc.CD_GUIA = pg.CD_GUIA_COBRANCA
		join itens_agrupados ia on
			pg.cd_procedimento = ia.cod_item 
		where
			gc.DS_GUIA_STATUS = 'GUIA_VALIDA'
			and pg.dt_execucao >= '|start|'
			and pg.dt_execucao < '|end|'
			and upper(grupo) in ('SAD', 'HONORARIO MEDICO CIRURGICO')
	),
	anatomos as (
		select distinct
			gc.cd_guia as an_cd_guia, pg.cd_procedimento as an_cd_procedimento, pg.dt_execucao as an_dt_execucao
		from
			guia_cobranca gc
		join procedimento_guia pg on
			gc.CD_GUIA = pg.CD_GUIA_COBRANCA
		where
			gc.DS_GUIA_STATUS = 'GUIA_VALIDA'
			and pg.dt_execucao >= '|start|'
			and pg.dt_execucao < '|end|'
			and pg.cd_procedimento = %d 
	)
select
	ia.ia_cd_procedimento,
	DATE_PART('day', AGE(ia.ia_dt_execucao, an.an_dt_execucao)) as dias, 
	count(distinct gc.cd_guia) as conta
from
	guia_cobranca gc
join procedimento_guia pg on
	gc.CD_GUIA = pg.CD_GUIA_COBRANCA
join item_agrupado ia on 
	ia.ia_cd_guia = gc.cd_guia 
join anatomos an on 
	an.an_cd_guia = gc.cd_guia 
where
	gc.DS_GUIA_STATUS = 'GUIA_VALIDA'
	and ia.ia_dt_execucao >= an.an_dt_execucao
	and an.an_dt_execucao <= ia.ia_dt_execucao + interval '|days|' day
	and pg.dt_execucao >= '|start|'
	and pg.dt_execucao < '|end|'
	and pg.cd_procedimento = %d
	and (pg.cd_procedimento = ia.ia_cd_procedimento or pg.cd_procedimento = an.an_cd_procedimento) 
group by 
	ia.ia_cd_procedimento, 
	DATE_PART('day', AGE(ia.ia_dt_execucao, an.an_dt_execucao)) 
having 
	DATE_PART('day', AGE(ia.ia_dt_execucao, an.an_dt_execucao)) <= |days|
order by 2, 1, 3
