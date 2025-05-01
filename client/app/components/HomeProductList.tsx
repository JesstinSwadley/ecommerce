import React from 'react'

type Product = {
	id: number,
	name: string 
}

const HomeProductList = async () => {
	const res = await fetch(`${process.env.GO_API}/products`);

	const products: Product[] = await res.json();

	return (
		<div>
			{products.map(product => <li key={product.id}><span>{product.name}</span></li>)}
		</div>
	)
}

export default HomeProductList