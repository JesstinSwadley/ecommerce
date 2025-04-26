import React from 'react'

const HomeProductList = () => {
	const Products = [
		{
			name: "Product 1",
			id: 1
		},
		{
			name: "Product 2",
			id: 2
		}
	]

	return (
		<div>
			{Products.map(product => <li key={product.id}><span>{product.name}</span></li>)}
		</div>
	)
}

export default HomeProductList