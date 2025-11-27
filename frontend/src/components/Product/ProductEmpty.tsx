export const ProductEmpty = () => {
	return (
		<div className="rounded-lg border-4 border-border border-dashed bg-muted/20 p-12 text-center">
			<p className="font-bold font-mono text-muted-foreground text-xl">
				No ingredients found in the provided text.
			</p>
		</div>
	);
};
