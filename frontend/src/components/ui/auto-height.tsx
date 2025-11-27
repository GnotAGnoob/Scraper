import { type ReactNode, useEffect, useLayoutEffect, useRef, useState } from "react";
import { cn } from "@/lib/utils";

interface AutoHeightProps {
	children: ReactNode;
	className?: string;
	durationMs?: number;
	onHeightTransitionEnd?: () => void;
	deferUnmountOnShrink?: boolean;
}

export const AutoHeight = ({
	children,
	className,
	durationMs = 300,
	onHeightTransitionEnd,
	deferUnmountOnShrink = true,
}: AutoHeightProps) => {
	const containerRef = useRef<HTMLDivElement | null>(null);
	const contentRef = useRef<HTMLDivElement | null>(null);
	const isFirstRef = useRef(true);
	const prevChildrenRef = useRef<ReactNode>(children);
	const prevHeightRef = useRef(0);
	const [ghost, setGhost] = useState<ReactNode | null>(null);
	const [height, setHeight] = useState(0);

	useEffect(() => {
		const container = containerRef.current;
		const content = contentRef.current;
		if (!container || !content) return;

		const handleTransitionEnd = (e: TransitionEvent) => {
			if (e.propertyName !== "height") return;

			requestAnimationFrame(() => {
				setGhost(null);
				onHeightTransitionEnd?.();
			});
		};
		container.addEventListener("transitionend", handleTransitionEnd);

		let raf1: number | null = null;
		let raf2: number | null = null;

		const observer = new ResizeObserver((entries) => {
			for (const entry of entries) {
				const newHeight = Math.ceil(entry.contentRect.height);

				if (isFirstRef.current) {
					isFirstRef.current = false;
					setHeight(newHeight);
					prevHeightRef.current = newHeight;
					continue;
				}

				const from = prevHeightRef.current;
				const to = newHeight;
				if (from === to) continue;

				if (deferUnmountOnShrink && to < from && ghost === null) {
					setGhost(prevChildrenRef.current);
				}

				if (raf1) cancelAnimationFrame(raf1);
				if (raf2) cancelAnimationFrame(raf2);
				raf1 = requestAnimationFrame(() => {
					raf2 = requestAnimationFrame(() => {
						setHeight(to);
						prevHeightRef.current = to;
					});
				});

				prevChildrenRef.current = children;
			}
		});

		observer.observe(content);
		return () => {
			if (raf1) cancelAnimationFrame(raf1);
			if (raf2) cancelAnimationFrame(raf2);
			observer.disconnect();
			container.removeEventListener("transitionend", handleTransitionEnd);
		};
	}, [ghost, deferUnmountOnShrink, onHeightTransitionEnd, children]);

	// biome-ignore lint/correctness/useExhaustiveDependencies: we want to run this when the content changes
	useLayoutEffect(() => {
		const content = contentRef.current;
		if (!deferUnmountOnShrink || !content) return;

		const measured = Math.ceil(content.getBoundingClientRect().height);

		let raf1: number | null = null;
		let raf2: number | null = null;

		if (height && measured < height && ghost === null) {
			setGhost(prevChildrenRef.current);

			if (raf1) cancelAnimationFrame(raf1);
			if (raf2) cancelAnimationFrame(raf2);
			raf1 = requestAnimationFrame(() => {
				raf2 = requestAnimationFrame(() => {
					setHeight(measured);
					prevHeightRef.current = measured;
				});
			});
		}

		return () => {
			if (raf1) cancelAnimationFrame(raf1);
			if (raf2) cancelAnimationFrame(raf2);
		};
	}, [children, deferUnmountOnShrink, height, ghost]);

	return (
		<div
			ref={containerRef}
			className={cn("overflow-hidden transition-[height] ease-out will-change-[height]", className)}
			style={{
				height: !height ? undefined : `${height}px`,
				transitionDuration: `${durationMs}ms`,
			}}
		>
			<div ref={contentRef} className={cn("relative will-change-contents")}>
				<div className={cn(ghost && "opacity-0")}>{children}</div>
				{ghost ? (
					<div className={cn("pointer-events-none absolute inset-0 z-10")} aria-hidden>
						{ghost}
					</div>
				) : null}
			</div>
		</div>
	);
};
