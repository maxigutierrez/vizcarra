import React from 'react';
import { connect } from 'react-redux';
import './styles.scss'
import { message, Table, Breadcrumb, Icon, Input, Button, Dropdown, Menu, Modal, Row, Col } from 'antd';
import QueueAnim from 'rc-queue-anim';
import queryString from 'query-string';
import { api } from './api';
import Nuevo from './Nuevo';
import Editar from './Editar';

class Productos extends React.Component {

	constructor(props) {
		super(props)
		this.state = {
			q: '',
			searchText: '',
			data: [],
			pagination: { showSizeChanger: true, pageSizeOptions: ['10', '25', '50', '100'] },
			loading: false,
			columns: [
				{
					title: 'Id',
					dataIndex: 'id',
					sorter: true,
					key: 'id',
					width: 100,
				},
				{
					title: 'Producto',
					dataIndex: 'producto',
					sorter: true,
					key: 'producto',
					width: 300,
				},
				{
					title: 'Acción',
					key: 'action',
					align: 'right',
					width: 1,
					render: (text, record) => {
						return (
							<Dropdown trigger={['click']} overlay={
								<Menu>
									<Menu.Item key="1" onClick={() => this.setState({ openEditar: true, registro: record })}>Editar</Menu.Item>
									<Menu.Item key="2" onClick={() => this.showDeleteConfirm(record.id)}>Borrar</Menu.Item>
								</Menu>
							}>
								<Button size="small">
									<Icon type="ellipsis" />
								</Button>
							</Dropdown>
						)
					},
				}],
		}
	}
	componentDidMount() {
		this.fetch();
	}

	showDeleteConfirm = async (id) => {
		const _this = this;
		Modal.confirm({
			title: '¿Esta seguro que desea borrar?',
			okText: 'Aceptar',
			okType: 'danger',
			cancelText: 'Cancelar',
			async onOk() {
				const response = await api.productos.delete(id);
				if (response.status !== "success") {
					message.error(response.message, 7);
				}
				_this.fetch();
			},
			onCancel() {
				console.log('Cancel');
			},
		});
	}

	handleTableChange = (pagination, filters, sorter) => {
		const pager = { ...this.state.pagination };
		pager.current = pagination.current;
		this.setState({
			pagination: pager,
		});
		this.fetch({
			limit: pagination.pageSize,
			page: pagination.current,
			sortField: sorter.field,
			sortOrder: sorter.order === "ascend" ? "ASC" : "DESC",
			...filters,
		});
	}

	fetch = async (params = {}) => {

		this.setState({ loading: true });

		try {
			const response = await api.productos.getAll(queryString.stringify({
				...params,
				q: this.state.q ? this.state.q : '',
			}))

			const pagination = { ...this.state.pagination };

			if (response.status === "success") {
				pagination.total = response.data.totalDataSize;

				this.setState({
					loading: false,
					data: response.data.productos,
					pagination,
				});
			} else {
				this.setState({
					loading: false,
					data: [],
					pagination,
				});
				message.error(response.message, 5);
			}
		} catch (e) {
			this.setState({
				loading: false,
				data: [],
			});
			message.error(e.toString(), 5);
		}
	}


	search = (text) => {

		this.setState({ q: text }, () => {
			this.fetch();
		})

	}

	render() {
		let columns = [...this.state.columns];
		return (
			<div className="container-fluid no-breadcrumb">
				<QueueAnim type="bottom" className="ui-animate">
					<Breadcrumb>
						<Breadcrumb.Item>Inicio</Breadcrumb.Item>
						<Breadcrumb.Item>Productos</Breadcrumb.Item>
					</Breadcrumb>
					<div className="box box-default box-ant-table-v1" style={{ marginTop: 15 }}>
						<div className="box-body">
							<div className="row pb-3">
								<div className="col-md-6">
									<h4>Productos</h4>
								</div>
								<div className="col-md-6 d-flex justify-content-end">
									<Button
										type="primary"
										icon="plus"
										onClick={() => this.setState({ openNuevo: true })}
									>
										Nuevo
									</Button>
								</div>
							</div>
							<Row >
								<Col style={{ float: 'right' }} sm={{ span: 8 }} xs={{ span: 24 }}>
									<Input.Search
										placeholder="Buscar..."
										onSearch={value => this.search(value)}
										style={{ marginBottom: 10, marginTop: 0, float: 'right', display: 'inline-block' }}
									/>
								</Col>
							</Row>
							<Table
								bordered={false}
								columns={columns}
								rowKey={record => record.id}
								dataSource={this.state.data}
								pagination={this.state.pagination}
								loading={this.state.loading}
								onChange={this.handleTableChange}
								onRow={(record, rowIndex) => {
									return {
										onDoubleClick: (e) => {
											this.setState({ openEditar: true, registro: record, })
										},
									};
								}}
							/>
						</div>
					</div>
				</QueueAnim>
				{this.state.openNuevo &&

					<Nuevo
						closeModal={() => {
							this.setState({ openNuevo: false });
							this.fetch();
						}}
					/>
				}
				{this.state.openEditar &&
					<Editar
						// cargarDatos= {(e)=>{
						// 		let newdata ={
						// 		productos:'Cheque',
						// 		banco:e.banco,
						// 		};
						// 		// this.state.data.push(newdata);
						// 		console.log(newdata)
						// 		}
						// 	}
						closeModal={() => {
							this.setState({ openEditar: false });
							this.fetch();
						}}
						data={this.state.registro}
					/>
				}
			</div>
		);
	}
}

const mapStateToProps = (state, ownProps) => {
	return {
		user: state.user,
	}
};

export default connect(
	mapStateToProps
)(Productos);
