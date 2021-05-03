
		import React from 'react';
		import {api} from './api';
		import { connect } from 'react-redux';
		import { message, Modal, Form, Input, Row, Col } from 'antd';
		import './styles.scss'
		
		const FormItem = Form.Item;
		
		class Editar extends React.Component {
		
			state = {
				disabledLogin: false,
				confirmLoading: false,
			}
		
			async componentDidMount(){
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
							}else {
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
							<h4 style={{marginBottom:15}}>Editar producto</h4>
							<Form style={{marginTop:10}}>
								<Row gutter={16}>
									<Col span={24}>
										<FormItem label="Producto" {...{
											labelCol: {sm: { span: 5 },},
											wrapperCol: {sm: { span: 19 },},
										}}>
											{getFieldDecorator('producto', {
												initialValue: this.props.data.producto ? this.props.data.producto : '',
												rules: [{ required: true, message: ' ' }],
												})(
												<Input size="default" />
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
		