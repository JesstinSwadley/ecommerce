'use client';
import React from 'react'

const CreateNewProductForm = () => {
	return (
		<>
			<form action="">
				<div>
					<label 
						htmlFor="productNameInput">
							Poll Query
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