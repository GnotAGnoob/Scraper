"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import { cn } from "@/lib/utils";
import { Popover, PopoverContent, PopoverTrigger } from "./popover";

type TruncatedTextProps = {
	text: string;
	className?: string;
	truncatedClassName?: string;
	maxLines?: number;
};

export const TruncatedText = ({ text, className, truncatedClassName, maxLines = 1 }: TruncatedTextProps) => {
	const [isTruncated, setIsTruncated] = useState(false);
	const textRef = useRef<HTMLButtonElement | null>(null);

	const isSingleLine = maxLines === 1;

	const measureTruncation = useCallback(() => {
		const el = textRef.current;
		if (!el) return;
		if (isSingleLine) {
			setIsTruncated(el.scrollWidth > el.clientWidth + 1);
			return;
		}
		const computed = window.getComputedStyle(el);
		const clone = document.createElement("div");
		clone.textContent = text;
		clone.style.position = "absolute";
		clone.style.visibility = "hidden";
		clone.style.pointerEvents = "none";
		clone.style.zIndex = "-1";
		clone.style.width = `${el.clientWidth}px`;
		clone.style.whiteSpace = computed.whiteSpace;
		clone.style.font = computed.font;
		clone.style.fontSize = computed.fontSize;
		clone.style.fontWeight = computed.fontWeight;
		clone.style.lineHeight = computed.lineHeight;
		clone.style.letterSpacing = computed.letterSpacing;
		clone.style.wordBreak = computed.wordBreak;
		clone.style.wordWrap = computed.wordWrap;
		clone.style.overflow = "visible";
		document.body.appendChild(clone);
		const fullHeight = clone.scrollHeight;
		clone.remove();
		setIsTruncated(fullHeight > el.clientHeight + 1);
	}, [isSingleLine, text]);

	useEffect(() => {
		measureTruncation();
		if (!textRef.current) return;
		const ro = new ResizeObserver(() => {
			measureTruncation();
		});
		ro.observe(textRef.current);
		return () => {
			ro.disconnect();
		};
	}, [measureTruncation]);

	return (
		<Popover isHoverEnabled={isTruncated}>
			<PopoverTrigger asChild>
				<button
					ref={textRef}
					type="button"
					className={cn(
						"focus-visible:outline-none",
						isSingleLine ? "truncate" : "overflow-hidden",
						className,
						isTruncated && truncatedClassName,
					)}
					style={
						!isSingleLine
							? {
									WebkitLineClamp: maxLines,
									display: "-webkit-box",
									WebkitBoxOrient: "vertical",
								}
							: undefined
					}
					aria-label={text}
					tabIndex={0}
				>
					{text}
				</button>
			</PopoverTrigger>
			<PopoverContent align="center" side="top" className="w-96">
				<div className="wrap-break-word whitespace-pre-wrap text-sm leading-snug">{text}</div>
			</PopoverContent>
		</Popover>
	);
};

export default TruncatedText;
