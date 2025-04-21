'use client';
import React from 'react'

const CreateNewProductForm = () => {
	const NewProductFormAction = async (formData : FormData) => {
		const productName = formData.get("productNameInput");

		await fetch(`${process.env.GO_API}/products`, {
			headers: {
				'Content-Type': 'application/json',
			},
			method: 'POST',
			body: JSON.stringify({
				productName
			})
		});
	}

	return (
		<>
			<form 
				action={NewProductFormAction}>
				<div>
					<label 
						htmlFor="productNameInput">
							Product Name
					</label>
					<input
						id="productNameInput"
						type="text" />
				</div>
				<button 
					type="submit">
						Create New Product
				</button>
			</form>
		</>
	)
}

export default CreateNewProductForm