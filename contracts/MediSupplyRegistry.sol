// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title MediSupplyRegistry
 * @dev Smart contract para registrar y verificar hashes de transacciones médicas en blockchain
 * Este contrato permite registrar hashes junto con CIDs de IPFS para trazabilidad
 */
contract MediSupplyRegistry {
    // Estructura para almacenar información de registro
    struct Registro {
        bytes32 hash;
        string cid;
        address registrador;
        uint256 timestamp;
        bool existe;
    }

    // Mapeo de hash de transacción -> Registro
    mapping(bytes32 => Registro) public registros;
    
    // Mapeo para rastrear registros por cuenta
    mapping(address => bytes32[]) public registrosPorCuenta;
    
    // Array para mantener lista de todos los registros
    bytes32[] public todosLosRegistros;
    
    // Eventos para logging
    event HashRegistrado(
        bytes32 indexed hashTransaccion,
        bytes32 indexed hash,
        string cid,
        address indexed registrador,
        uint256 timestamp
    );
    
    event HashVerificado(
        bytes32 indexed hashTransaccion,
        bytes32 indexed hash,
        bool valido
    );

    /**
     * @dev Registra un hash de transacción junto con su CID de IPFS
     * @param hashTransaccion El hash único de la transacción
     * @param hash El hash de los datos médicos
     * @param cid El CID de IPFS donde están almacenados los datos
     * @return hashTx Hash de la transacción de registro
     */
    function registrarHash(
        bytes32 hashTransaccion,
        bytes32 hash,
        string memory cid
    ) public returns (bytes32 hashTx) {
        // Verificar que el hash no esté ya registrado
        require(
            !registros[hashTransaccion].existe,
            "El hash ya está registrado"
        );

        // Crear nuevo registro
        Registro memory nuevoRegistro = Registro({
            hash: hash,
            cid: cid,
            registrador: msg.sender,
            timestamp: block.timestamp,
            existe: true
        });

        // Guardar registro
        registros[hashTransaccion] = nuevoRegistro;
        registrosPorCuenta[msg.sender].push(hashTransaccion);
        todosLosRegistros.push(hashTransaccion);

        // Emitir evento
        emit HashRegistrado(
            hashTransaccion,
            hash,
            cid,
            msg.sender,
            block.timestamp
        );

        return hashTransaccion;
    }

    /**
     * @dev Verifica si un hash está registrado y es válido
     * @param hashTransaccion El hash de la transacción a verificar
     * @param hashEsperado El hash esperado de los datos
     * @return valido True si el hash está registrado y coincide
     */
    function verificarHash(
        bytes32 hashTransaccion,
        bytes32 hashEsperado
    ) public view returns (bool valido) {
        Registro memory registro = registros[hashTransaccion];
        
        if (!registro.existe) {
            emit HashVerificado(hashTransaccion, bytes32(0), false);
            return false;
        }

        bool esValido = registro.hash == hashEsperado;
        
        return esValido;
    }

    /**
     * @dev Obtiene información completa de un registro
     * @param hashTransaccion El hash de la transacción
     * @return hash El hash de los datos
     * @return cid El CID de IPFS
     * @return registrador La dirección que registró
     * @return timestamp El timestamp del registro
     * @return existe Si el registro existe
     */
    function obtenerRegistro(
        bytes32 hashTransaccion
    )
        public
        view
        returns (
            bytes32 hash,
            string memory cid,
            address registrador,
            uint256 timestamp,
            bool existe
        )
    {
        Registro memory registro = registros[hashTransaccion];
        return (
            registro.hash,
            registro.cid,
            registro.registrador,
            registro.timestamp,
            registro.existe
        );
    }

    /**
     * @dev Obtiene todos los registros de una cuenta
     * @param cuenta La dirección de la cuenta
     * @return listaHashes Array de hashes registrados por la cuenta
     */
    function obtenerRegistrosPorCuenta(
        address cuenta
    ) public view returns (bytes32[] memory listaHashes) {
        return registrosPorCuenta[cuenta];
    }

    /**
     * @dev Obtiene el número total de registros
     * @return total El número total de registros
     */
    function totalRegistros() public view returns (uint256 total) {
        return todosLosRegistros.length;
    }

    /**
     * @dev Obtiene un registro por índice
     * @param indice El índice en la lista de registros
     * @return hashTransaccion El hash de la transacción
     */
    function obtenerRegistroPorIndice(
        uint256 indice
    ) public view returns (bytes32 hashTransaccion) {
        require(indice < todosLosRegistros.length, "Índice fuera de rango");
        return todosLosRegistros[indice];
    }
}

