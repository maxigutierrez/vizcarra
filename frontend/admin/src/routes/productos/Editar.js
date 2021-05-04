
import React from 'react';
import { api } from './api';
import { connect } from 'react-redux';
import { message, Modal, Form, Input, Row, Col, Select } from 'antd';
import './styles.scss'

const FormItem = Form.Item;
const Option = Select.Option;

class Editar extends React.Component {

	state = {
		disabledLogin: false,
		confirmLoading: false,
	}

	async componentDidMount() {
		const response = await api.marcas.getAll()
		if (response.status === "success") {
			this.setState({ marcas: response.data.marcas, });
		} else {
			message.error(response.message, 5);
		}
	}

	handleSubmit = (e) => {
		e.preventDefault();

		this.props.form.validateFields(async (err, values) => {
			if (!err) {
				try {
					this.setState({
						confirmLoading: true,
					})
					const response = await api.productos.update({
						...values,
						id: this.props.data.id,
					});
					if (response.status === "success") {
						this.props.closeModal();
					} else {
						message.error(response.message, 7);
					}
				} catch (e) {
					message.error(e.toString(), 7);
				} finally {
					this.setState({
						confirmLoading: false,
					})
				}
			}
		})
	}

	render() {
		console.log(this.props)
		const { getFieldDecorator } = this.props.form;
		return (
			<Modal
				visible={true}
				confirmLoading={this.state.confirmLoading}
				onOk={this.handleSubmit}
				onCancel={this.props.closeModal}
				cancelText="Cancelar"
				okText="Guardar"
				width={600}
			>
				<section className="form-v1-container col-md-12">
					<h4 style={{ marginBottom: 15 }}>Editar producto</h4>
					<Form style={{ marginTop: 10 }}>
						<Row gutter={16}>
							<Col span={24}>
								<FormItem label="Producto" {...{
									labelCol: { sm: { span: 5 }, },
									wrapperCol: { sm: { span: 19 }, },
								}}>
									{getFieldDecorator('producto', {
										initialValue: this.props.data ? this.props.data.producto : '',
										rules: [{ required: true, message: ' ' }],
									})(
										<Input size="default" />
									)}
								</FormItem>
							</Col>
						</Row>
						<Row gutter={16}>
							<Col span={24}>
								<FormItem label="Marca" {...{
									labelCol: { sm: { span: 5 }, },
									wrapperCol: { sm: { span: 19 }, },
								}}>
									{getFieldDecorator('marcas_id', {
										initialValue: this.props.data ? this.props.data.marcas_id : '',
										rules: [{ required: true, message: 'Campo obligatorio' }],
									})(

										<Select
											placeholder="Seleccione una marca"
											optionFilterProp="children"
											showSearch
											filterOption={(input, option) =>
												option.props.children ? option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0 : true
											}
										>
											{this.state.marcas && this.state.marcas.map((data, index) => {
												return <Option value={data.id} key={index}>{data.marca}</Option>
											})}
										</Select>
									)}
								</FormItem>
							</Col>
						</Row>
					</Form>
				</section>
			</Modal>
		);
	}
}

const mapStateToProps = (state) => ({
	user: state.user,
});

const WrappedEditar = Form.create()(Editar);

export default connect(
	mapStateToProps,
)(WrappedEditar);
