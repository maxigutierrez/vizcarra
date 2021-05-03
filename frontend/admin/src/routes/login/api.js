import 'isomorphic-fetch';
import {config} from 'config'

export const api = {
    auth:{
    //     async login(data) {
    //         const response = await fetch(config.URL_API + '/login', {
	// 		  method: 'POST',
	// 		  headers: {
	// 		    'Content-Type': 'application/json'
	// 		  },
	// 		  body: JSON.stringify(data)
	// 		});
    //         const responseData = await response.json();
    //         return responseData;
    //     },
    //     async getAuthenticatedUser() {
    //         const response = await fetch(config.URL_API + '/authenticated', {
    //             method: 'GET',
    //             headers: {'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`},
	// 		});
    //         const data = await response.json();
    //         return data;
    //     },
    // }
};