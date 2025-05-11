import React from 'react'

const UpdateProductForm = () => {
	return (
		<>
			<form action="">
				<div>
					<label htmlFor="">Product Id</label>
					<input type="text" />
				</div>
				<div>
					<label htmlFor="">Update Product Name</label>
					<input type="text" />
				</div>

				<button
					type="submit">Update Product</button>
			</form>
		</>
	)
}

export default UpdateProductForm