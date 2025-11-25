interface OverlayProps {
	isOpen: boolean;
	onClose: () => void;
}

export const Overlay = ({ isOpen, onClose }: OverlayProps) => {
	return (
		<button
			type="button"
			className={`fixed inset-0 bg-overlay z-40 w-full h-full cursor-default transition-opacity duration-300 ${
				isOpen ? "opacity-100 visible" : "opacity-0 pointer-events-none invisible"
			}`}
			onClick={onClose}
			onKeyDown={(e) => {
				if (e.key === "Escape") onClose();
			}}
			aria-label="Close menu"
		/>
	);
};
