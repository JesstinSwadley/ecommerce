import React from 'react'
import ProductCard from './ProductCard';

type Product = {
	id: number,
	name: string 
}

const HomeProductList = async () => {
	const res = await fetch(`${process.env.GO_API}/products`);

	const products: Product[] = await res.json();

	return (
		<div>
			{products.map(product => <ProductCard key={product.id} name={product.name} />)}
		</div>
	)
}

export default HomeProductList