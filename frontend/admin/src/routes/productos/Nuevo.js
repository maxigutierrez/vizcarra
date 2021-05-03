
import React from 'react';
import { api } from './api';
import { connect } from 'react-redux';
import { message, Modal, Form, Input, Row, Col } from 'antd';
import './styles.scss'

const FormItem = Form.Item;

class Nuevo extends React.Component {

	state = {
		disabledLogin: false,
		confirmLoading: false,
	}

	async componentDidMount() {
	}

	handleSubmit = (e) => {
		e.preventDefault();

		this.props.form.validateFields(async (err, values) => {
			if (!err) {
				try {
					this.setState({
						confirmLoading: true,
					})
					const response = await api.productos.create({
						...values,
					});
					if (response.status === "success") {

					} else {
						message.error(response.message, 7);
					}
				} catch (e) {
					message.error(e.toString(), 7);
					console.log(e)
				} finally {
					this.setState({
						confirmLoading: false,
					})
					this.props.closeModal();
				}
			}
		})
	}

	render() {
		const { getFieldDecorator } = this.props.form;
		return (
			<Modal
				visible={true}
				confirmLoading={this.state.confirmLoading}
				onOk={this.handleSubmit}
				onCancel={this.props.closeModal}
				cancelText="Cancelar"
				okText="Crear"
				width={600}
			>
				<section className="form-v1-container col-md-12">
					<h4 style={{ marginBottom: 15 }}>Nuevo Producto</h4>
					<Form style={{ marginTop: 10 }}>
						<Row gutter={16}>
							<Row gutter={16}>
								<Col span={24}>
									<FormItem label="Producto" {...{
										labelCol: { sm: { span: 5 }, },
										wrapperCol: { sm: { span: 16 }, },
									}}>
										{getFieldDecorator('producto', {
											rules: [{ required: true, message: 'Campo obligatorio' }],
										})(
											<Input size="default"/>
										)}
									</FormItem>
								</Col>
							</Row>
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

const WrappedNuevo = Form.create()(Nuevo);

export default connect(
	mapStateToProps,
)(WrappedNuevo);


