
import 'isomorphic-fetch';
import { config } from 'config'

export const api = {
	marcas: {
		async create(data) {
			const response = await fetch(config.URL_API + '/marcas', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					// 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
				},
				body: JSON.stringify(data)
			});
			const responseData = await response.json();
			return responseData;
		},
		async getAll(query = '') {
			const response = await fetch(config.URL_API + '/marcas?' + query, {
				method: 'GET',
				// headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
			});
			const data = await response.json();
			return data;
		},
		async get(id) {
			const response = await fetch(config.URL_API + `/marcas/${id}`, {
				method: 'GET',
				headers: { 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}` },
			});
			const data = await response.json();
			return data;
		},
		async update(data) {
			const response = await fetch(config.URL_API + '/marcas', {
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
		async delete(id) {
			const response = await fetch(config.URL_API + '/marcas/' + id, {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					// 'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`,
				},
			});
			const responseData = await response.json();
			return responseData;
		},
	},
};
