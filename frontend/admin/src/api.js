import 'isomorphic-fetch';
import { config } from './config'

export const api = {
    auth: {
        // async login(data) {
        //     const response = await fetch(config.URL_API + '/login', {
        //         method: 'POST',
        //         headers: {
        //             'Content-Type': 'application/json'
        //         },
        //         body: JSON.stringify(data)
        //     });
        //     const responseData = await response.json();
        //     return responseData;
        // },
        // async getAuthenticatedUser() {
        //     const response = await fetch(config.URL_API + '/authenticated', {
        //         method: 'GET',
        //         headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
        //     });
        //     const data = await response.json();
        //     return data;
        // },
    },
    provincias: {
        async getAll() {
            const response = await fetch(config.URL_API + `/provincias`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    localidades: {
        async getAll(params = '') {
            const response = await fetch(config.URL_API + '/clientes/localidades?' + params, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    direcciones: {
        async getCalles(params = '') {
            const response = await fetch(config.URL_API_LOCALIDAD + '/calles?' + params , {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await response.json();
            return data;
        },
        async getDirecciones(params = '') {
            const response = await fetch(config.URL_API_LOCALIDAD + '/direcciones?' + params + '&aplanar', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await response.json();
            return data;
        },
    },
    localidades1: {
        async getAll(params = '') {
            const response = await fetch(config.URL_API_LOCALIDAD + '/localidades?' + params, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                     'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const data = await response.json();
            return data;
        },
        async get(id) {
            const response = await fetch(config.URL_API + `/localidad/${id}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const data = await response.json();
            return data;
        },
    },
    domicilios_tipos: {
        async getAll(params = '') {
            const response = await fetch(config.URL_API + '/domicilios_tipos?' + params, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
                },
            });
            const data = await response.json();
            return data;
        },
    },
    franquicias: {
        async getAllconNodos(query = '') {
            const response = await fetch(config.URL_API + `/franquicias/all?` + query, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async getAllDominios(query = '') {
            const response = await fetch(config.URL_API + `/franquicias/dominios?` + query, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    user: {
        async create(data) {
            const response = await fetch(config.URL_API + '/usuarios', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
        async getAll(query = '') {
            const response = await fetch(config.URL_API + `/usuarios?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async get(id) {
            const response = await fetch(config.URL_API + `/usuarios/${id}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async update(data) {
            const response = await fetch(config.URL_API + '/usuarios', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
        async baja(id) {
            const response = await fetch(config.URL_API + '/usuarios/baja/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const responseData = await response.json();
            return responseData;
        },
        async alta(id) {
            const response = await fetch(config.URL_API + '/usuarios/alta/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const responseData = await response.json();
            return responseData;
        },
        async getAllFranquicias() {
            const response = await fetch(config.URL_API + `/usuarios/franquicias`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const responseData = await response.json();
            return responseData;
        },
        async getAllRoles() {
            const response = await fetch(config.URL_API + '/usuarios/roles', {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async getAllSistemas() {
            const response = await fetch(config.URL_API + '/usuarios/sistemas', {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async changePassword(id, data) {
            const response = await fetch(config.URL_API + '/password/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
        async blanquearClave(id, data) {
            const response = await fetch(config.URL_API + '/blanqueo/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        }
    },
    afip: {
        async consultarCuit(cuit = '') {
            const response = await fetch(`http://10.8.1.9:5600/afip/cuit/${cuit}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    google_api: {
        async autocomplete(text = '') {
            const response = await fetch(`${config.URL_SOCKET_API}/google/autocomplete?q=${text}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await response.json();
            return data;
        },
        async getByPlaceID(placeid = '') {
            const response = await fetch(`${config.URL_SOCKET_API}/google/place/details?placeid=${placeid}`, {
                method: 'GET',
            });
            const data = await response.json();
            return data;
        },
    },
    nominatim: {
        async search(q) {
            const response = await fetch(`https://nominatim.openstreetmap.org/search?q=${q}&format=json&addressdetails=1&polygon_geojson=1&country=Argentina`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await response.json();
            return data;
        },
    },
    graphhopper: {
        async optimize(data) {
            const response = await fetch('https://graphhopper.com/api/1/vrp/optimize?key=370d5360-c71b-4fba-a3fe-04ff2f3c8e07', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
        async getSolution(job_id) {
            const response = await fetch(`https://graphhopper.com/api/1/vrp/solution/${job_id}?key=370d5360-c71b-4fba-a3fe-04ff2f3c8e07`, {
                method: 'GET',
            });
            const data = await response.json();
            return data;
        }

    },
    retiros:{
        async getCount(query = '') {
            const response = await fetch(config.URL_API + '/retiros/count' , {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },    
    },
    autorizaciones: {
        async getAll(query = '') {
            const response = await fetch(config.URL_API + '/autorizaciones/alertas?' + query, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async aprobar(id) {
            const response = await fetch(config.URL_API + '/autorizaciones/aprobar/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const responseData = await response.json();
            return responseData;
        },
        async rechazar(id) {
            const response = await fetch(config.URL_API + '/autorizaciones/rechazar/' + id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const responseData = await response.json();
            return responseData;
        },
    },
    comprobantes: {
        async getComprobante(id) {
            const response = await fetch(config.URL_API + '/comprobantes/numero/' + id, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
            });
            const responseData = await response.json();
            return responseData;
        },
    },
    comprobantes_tipos:{
        async create(data) {
            const response = await fetch(config.URL_API + '/comprobantestipos', {
			  method: 'POST',
			  headers: {
			    'Content-Type': 'application/json',
                'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
			  },
			  body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async getAll(query = '') {
            const response = await fetch(config.URL_API + `/comprobantestipos?` + query, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async get(id) {
            const response = await fetch(config.URL_API + `/comprobantestipos/${id}`, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async update(data){
            const response = await fetch(config.URL_API + '/comprobantestipos', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
                body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async delete(id){
            const response = await fetch(config.URL_API + '/comprobantestipos/'+id, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
			});
            const responseData = await response.json();
            return responseData;
        },
    },
    condiciones_iva: {
        async getAll(query = '') {
            const response = await fetch(config.URL_API + `/condicioniva?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    cobertura: {
        async getCobertura(data) {
            const response = await fetch(config.URL_API + '/cobertura', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
    },
    documentos_tipos: {
        async getAll() {
            const response = await fetch(config.URL_API + `/documentotipos`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    clientes:{
        async create(data) {
            const response = await fetch(config.URL_API + '/clientes', {
			  method: 'POST',
			  headers: {
			    'Content-Type': 'application/json',
                'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
			  },
			  body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async getAll(query = '') {
            const response = await fetch(config.URL_API + `/clientes?${query}`, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async get(id) {
            const response = await fetch(config.URL_API + `/clientes/${id}`, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async update(data){
            const response = await fetch(config.URL_API + '/clientes', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
                body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async baja(id){
            const response = await fetch(config.URL_API + '/clientes/baja/'+id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
			});
            const responseData = await response.json();
            return responseData;
        },
        async verificar(cuit) {
            const response = await fetch(config.URL_API + `/clientes_verificar/${cuit}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async validarDatos(query = ''){
            const response = await fetch(config.URL_API + `/clientes_validacion?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async alta(id){
            const response = await fetch(config.URL_API + '/clientes/alta/'+id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
			});
            const responseData = await response.json();
            return responseData;
        },
    },
    clientes_domicilios:{
        async create(data) {
            const response = await fetch(config.URL_API + '/clientes/domicilios', {
			  method: 'POST',
			  headers: {
			    'Content-Type': 'application/json',
                'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
			  },
			  body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async AsignarSucursal(franquicias_id,id) {
            console.log(franquicias_id, id)
            const response = await fetch(config.URL_API + `/clientes_domicilios/${franquicias_id}/${id}`, {
			  method: 'POST',
			  headers: {
			    'Content-Type': 'application/json',
                'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
			  },
			//   body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async getAll(id, query = '') {
            const response = await fetch(config.URL_API + `/clientes/${id}/domicilios?` + query, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async get(id) {
            const response = await fetch(config.URL_API + `/clientes/domicilios/${id}`, {
                method: 'GET',
                headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
            const data = await response.json();
            return data;
        },
        async update(data){
            const response = await fetch(config.URL_API + '/clientes/domicilios', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
                body: JSON.stringify(data)
			});
            const responseData = await response.json();
            return responseData;
        },
        async delete(id){
            const response = await fetch(config.URL_API + '/clientes/domicilios/'+id, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
                },
			});
            const responseData = await response.json();
            return responseData;
        },
    },
    coberturaGeo: {
        async get(query = '') {
            const response = await fetch(config.URL_API + `/coberturageo?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async getExcluido(query = '') {
            const response = await fetch(config.URL_API + `/coberturageoexluir?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
    agildata: {
        async login(data) {
            const response = await fetch(config.URL_AGILDATA + '/account/authenticate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });
            const responseData = await response.json();
            return responseData;
        },
        async getByDNI(query = '') {
            const response = await fetch(config.URL_API + `/getByDni?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async getByCUIT(cuit) {
            const response = await fetch(config.URL_API + `/getByCuit/${cuit}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
        async getExcluido(query = '') {
            const response = await fetch(config.URL_APIG + `/coberturageoexluir?${query}`, {
                method: 'GET',
                headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
            });
            const data = await response.json();
            return data;
        },
    },
};
