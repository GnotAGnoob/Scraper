interface OverlayProps {
	isOpen: boolean;
	onClose: () => void;
}

export const Overlay = ({ isOpen, onClose }: OverlayProps) => {
	return (
		<button
			type="button"
			className={`fixed inset-0 z-40 h-full w-full cursor-default bg-overlay transition-opacity duration-300 ${
				isOpen ? "visible opacity-100" : "pointer-events-none invisible opacity-0"
			}`}
			onClick={onClose}
			onKeyDown={(e) => {
				if (e.key === "Escape") onClose();
			}}
			aria-label="Close menu"
		/>
	);
};
