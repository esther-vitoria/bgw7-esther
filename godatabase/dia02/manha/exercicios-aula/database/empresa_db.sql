DROP DATABASE IF EXISTS empresa_db;
CREATE DATABASE empresa_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE empresa_db;

CREATE TABLE depto (
  depto_nro VARCHAR(255),
  depto_name VARCHAR(255),
  depto_locallity VARCHAR(255),
  PRIMARY KEY (depto_nro)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE employee (
  cod_emp VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  position VARCHAR(255) NOT NULL,
  start_date VARCHAR(255) NOT NULL,
  salary FLOAT NOT NULL,
  bonus FLOAT NOT NULL,
  depto_nro VARCHAR(255),
  PRIMARY KEY (cod_emp),
  FOREIGN KEY (depto_nro) REFERENCES depto(depto_nro)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- Inserir departamentos
INSERT INTO depto (depto_nro, depto_name, depto_locallity) VALUES
('D-000-1', 'Software', 'Los Tigres'),
('D-000-2', 'Sistemas', 'Guadalupe'),
('D-000-3', 'Contabilidade', 'La Roca'),
('D-000-4', 'Vendas', 'Plata');

-- Inserir funcionários
INSERT INTO employee (cod_emp, first_name, last_name, position, start_date, salary, bonus, depto_nro) VALUES
('E-0001', 'Cesar', 'Pinero', 'Vendedor', '2018-05-12', 8000.00, 1500.00, 'D-000-4'),
('E-0002', 'Yosep', 'Kowaleski', 'Analista', '2015-07-14', 14000.00, 0, 'D-000-2'),
('E-0003', 'Mariela', 'Barrios', 'Dieretor', '2014-05-06', 18500.00, 0, 'D-000-3'),
('E-0004', 'Jonathan', 'Aguilera', 'Vendedor', '2015-06-03', 8500.00, 1000.00, 'D-000-4'),
('E-0005', 'Daniel', 'Brezezicki', 'Vendedor', '2018-03-03', 8300.00, 1000.00, 'D-000-4'),
('E-0006', 'Mito', 'Barchuk', 'Presidente', '2014-06-05', 19000.00, 0, 'D-000-3'),
('E-0007', 'Emilio', 'Galarza', 'Desenvolvedor', '2014-08-02', 6000.00, 0, 'D-000-1');