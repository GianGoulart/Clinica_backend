CREATE TABLE `BD_ClinicaAbrao`.`pacientes` (
  `id` VARCHAR(26) NOT NULL,
  `cpf` VARCHAR(14) NULL,
  `nome` VARCHAR(128) NOT NULL,
  `telefone` VARCHAR(14) NULL,
  `convenio` VARCHAR(26) NULL,
  `plano` VARCHAR(45) NULL,
  `acomodacao` VARCHAR(14) NULL,
  `status` INT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `cpf_UNIQUE` (`cpf` ASC) VISIBLE);

CREATE TABLE `BD_ClinicaAbrao`.`medicos` (
  `id` VARCHAR(26) NOT NULL,
  `nome` VARCHAR(128) NOT NULL,
  `cpf` VARCHAR(14) NOT NULL,
  `banco_pf` VARCHAR(26) NULL,
  `agencia_pf` VARCHAR(6) NULL,
  `conta_pf` VARCHAR(14) NULL,
  `razao_social` VARCHAR(128) NULL,
  `banco_pj` VARCHAR(26) NULL,
  `agencia_pj` VARCHAR(6) NULL,
  `conta_pj` VARCHAR(14) NULL,
  `cnpj` VARCHAR(16) NULL,
  `status` INT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`));
  UNIQUE INDEX `cpf_UNIQUE` (`cpf` ASC) VISIBLE);

CREATE TABLE `BD_ClinicaAbrao`.`procedimentos` (
  `id` VARCHAR(26) NOT NULL,
  `id_paciente` VARCHAR(26) NOT NULL,
  `id_medico` VARCHAR(26) NOT NULL,
  `desc_procedimento` VARCHAR(128) NULL,
  `procedimento` INT(1) NOT NULL,
  `local_procedimento` INT(2) NOT NULL,
  `status` INT(1) NOT NULL,
  `data` BIGINT(20) NOT NULL,
  `valor` DECIMAL(14,2) NOT NULL,
  `esteira` INT(1) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `BD_ClinicaAbrao`.`comercial` (
  `id` VARCHAR(26) NOT NULL,
  `id_procedimento` VARCHAR(26) NOT NULL,
  `id_medico_part` VARCHAR(26) NULL,
  `funcao_medico_part` INT(1) NULL,
  `qtd_parcelas` INT(2) NULL,
  `valor_parcelas` DECIMAL(14,2) NULL,
  `tipo_pagamento` INT(1) NULL,
  `forma_pagamento` INT(1) NULL,
  `data_emissao_nf` BIGINT(20) NULL,
  `data_vencimento` BIGINT(20) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `BD_ClinicaAbrao`.`financeiro` (
  `id` VARCHAR(26) NOT NULL,
  `ìd_comercial` VARCHAR(26) NOT NULL,
  `data_pagamento` BIGINT(20) NULL,
  `data_compensacao` BIGINT(20) NULL,
  `plano_contas` INT(1) NULL,
  `conta` INT(2) NULL,
  `valor_ajuste` DECIMAL(14,2) NULL,
  `valor_liquido` DECIMAL(14,2) NULL,
  `obs` TEXT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `BD_ClinicaAbrao`.`acompanhamentos` (
  `id` VARCHAR(26) NOT NULL,
  `id_procedimento` VARCHAR(26) NOT NULL,
  `envio_protocolo` BIGINT(20) NULL,
  `solicitacao_previa` BIGINT(20) NULL,
  `confirmacao_solicitacao` BIGINT(20) NULL,
  `finalizacao_previa` BIGINT(20) NULL,
  `status_previa` INT NOT NULL,
  `envio_convenio` BIGINT(20) NULL,
  `liberacao` BIGINT(20) NULL,
  `repasse_paciente` BIGINT(20) NULL,
  `repasse_clinica` BIGINT(20) NULL,
  `status_reembolso` INT NOT NULL,
  `obs` TEXT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `BD_ClinicaAbrao`.`contas` (
  `id` VARCHAR(26) NOT NULL,
  `cod_conta` VARCHAR(10) NOT NULL,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`));


