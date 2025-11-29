"use client";

import * as PopoverPrimitive from "@radix-ui/react-popover";
import {
	type ComponentPropsWithoutRef,
	createContext,
	type ElementRef,
	forwardRef,
	type KeyboardEvent,
	type MouseEvent,
	type TouchEvent,
	useCallback,
	useContext,
	useEffect,
	useRef,
	useState,
} from "react";
import { cn } from "@/lib/utils";

export interface UsePopoverHoverOptions {
	isEnabled?: boolean;
	hoverDelayMs?: number;
	longPressMs?: number;
	initialOpen?: boolean;
}

interface UsePopoverHoverReturn {
	open: boolean;
	handleOpenChange: (isOpen: boolean) => void;
	triggerHandlers: PopoverHoverHandlers;
}

interface PopoverHoverContextValue {
	isHoverEnabled: boolean;
	handleOpenChange: (isOpen: boolean) => void;
	triggerHandlers: PopoverHoverHandlers | null;
}

const PopoverHoverContext = createContext<PopoverHoverContextValue>({
	isHoverEnabled: false,
	handleOpenChange: () => {},
	triggerHandlers: null,
});

interface PopoverExtendedProps extends Omit<UsePopoverHoverOptions, "isEnabled"> {
	isHoverEnabled?: boolean;
}

interface PopoverProps extends ComponentPropsWithoutRef<typeof PopoverPrimitive.Root>, PopoverExtendedProps {}

const Popover = ({
	isHoverEnabled = false,
	hoverDelayMs,
	longPressMs,
	initialOpen,
	open,
	onOpenChange,
	...props
}: PopoverProps) => {
	const hover = usePopoverHover({
		isEnabled: isHoverEnabled,
		hoverDelayMs,
		longPressMs,
		initialOpen,
	});

	const mergedOpen = open ?? hover.open;
	const mergedOnOpenChange = useCallback(
		(isOpen: boolean) => {
			if (isHoverEnabled) hover.handleOpenChange(isOpen);
			if (onOpenChange) onOpenChange(isOpen);
		},
		[isHoverEnabled, hover, onOpenChange],
	);

	return (
		<PopoverHoverContext.Provider
			value={{
				isHoverEnabled: isHoverEnabled,
				handleOpenChange: mergedOnOpenChange,
				triggerHandlers: isHoverEnabled ? hover.triggerHandlers : null,
			}}
		>
			<PopoverPrimitive.Root open={mergedOpen} onOpenChange={mergedOnOpenChange} {...props} />
		</PopoverHoverContext.Provider>
	);
};

type PopoverTriggerProps = ComponentPropsWithoutRef<typeof PopoverPrimitive.Trigger>;

const PopoverTrigger = forwardRef<ElementRef<typeof PopoverPrimitive.Trigger>, PopoverTriggerProps>(
	({ onMouseEnter, onMouseLeave, onTouchStart, onTouchEnd, onKeyDown, onClick, ...props }, ref) => {
		const ctx = useContext(PopoverHoverContext);

		const handleOnClick = useCallback(
			(event: MouseEvent<HTMLButtonElement>) => {
				event.stopPropagation();
				onClick?.(event);
			},
			[onClick],
		);

		const compose = useCallback(
			<E,>(userHandler: ((e: E) => void) | undefined, libHandler: ((e: E) => void) | undefined) => {
				return (e: E) => {
					if (userHandler) userHandler(e);
					if (ctx.isHoverEnabled && libHandler) libHandler(e);
				};
			},
			[ctx.isHoverEnabled],
		);

		return (
			<PopoverPrimitive.Trigger
				ref={ref}
				onMouseEnter={compose(onMouseEnter, ctx.triggerHandlers?.onMouseEnter)}
				onMouseLeave={compose(onMouseLeave, ctx.triggerHandlers?.onMouseLeave)}
				onTouchStart={compose(onTouchStart, ctx.triggerHandlers?.onTouchStart)}
				onTouchEnd={compose(onTouchEnd, ctx.triggerHandlers?.onTouchEnd)}
				onKeyDown={compose(onKeyDown, ctx.triggerHandlers?.onKeyDown)}
				onClick={handleOnClick}
				{...props}
			/>
		);
	},
);
PopoverTrigger.displayName = PopoverPrimitive.Trigger.displayName;

const PopoverContent = forwardRef<HTMLDivElement, ComponentPropsWithoutRef<typeof PopoverPrimitive.Content>>(
	({ className, align = "center", sideOffset = 4, collisionPadding = 8, ...props }, ref) => (
		<PopoverPrimitive.Portal>
			<PopoverPrimitive.Content
				ref={ref}
				align={align}
				sideOffset={sideOffset}
				collisionPadding={collisionPadding}
				className={cn(
					"data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95 z-50 w-72 rounded-(--radius) border bg-background p-4 text-popover-foreground shadow-lg outline-none data-[state=closed]:animate-out data-[state=open]:animate-in",
					className,
				)}
				{...props}
			/>
		</PopoverPrimitive.Portal>
	),
);
PopoverContent.displayName = PopoverPrimitive.Content.displayName;

export type PopoverHoverHandlers = {
	onMouseEnter: (e: MouseEvent<HTMLElement>) => void;
	onMouseLeave: (e: MouseEvent<HTMLElement>) => void;
	onTouchStart: (e: TouchEvent<HTMLElement>) => void;
	onTouchEnd: (e: TouchEvent<HTMLElement>) => void;
	onKeyDown: (e: KeyboardEvent<HTMLElement>) => void;
};

const DEFAULT_HOVER_OPEN_DELAY_MS = 200;
const DEFAULT_LONG_PRESS_MS = 400;

export const usePopoverHover = ({
	isEnabled: enabled = true,
	hoverDelayMs = DEFAULT_HOVER_OPEN_DELAY_MS,
	longPressMs = DEFAULT_LONG_PRESS_MS,
	initialOpen = false,
}: UsePopoverHoverOptions): UsePopoverHoverReturn => {
	const [open, setOpen] = useState(initialOpen);
	const hoverOpenTimer = useRef<number | null>(null);
	const longPressTimer = useRef<number | null>(null);

	const isTouchDevice =
		typeof window !== "undefined" && typeof matchMedia !== "undefined" && matchMedia("(hover: none)").matches;

	const clearHoverOpen = useCallback(() => {
		if (hoverOpenTimer.current) {
			window.clearTimeout(hoverOpenTimer.current);
			hoverOpenTimer.current = null;
		}
	}, []);

	const clearLongPress = useCallback(() => {
		if (longPressTimer.current) {
			window.clearTimeout(longPressTimer.current);
			longPressTimer.current = null;
		}
	}, []);

	const handleOpenChange = useCallback(
		(isOpen: boolean) => {
			if (!enabled) return;
			setOpen(isOpen);
		},
		[enabled],
	);

	const onMouseEnter = useCallback(() => {
		if (isTouchDevice || !enabled) return;

		clearHoverOpen();
		const id = window.setTimeout(() => setOpen(true), hoverDelayMs);
		hoverOpenTimer.current = id;
	}, [isTouchDevice, enabled, hoverDelayMs, clearHoverOpen]);

	const onMouseLeave = useCallback(() => {
		if (isTouchDevice || !enabled) return;

		clearHoverOpen();
		setOpen(false);
	}, [isTouchDevice, enabled, clearHoverOpen]);

	const onTouchStart = useCallback(() => {
		if (isTouchDevice || !enabled) return;

		clearLongPress();
		const id = window.setTimeout(() => setOpen(true), longPressMs);
		longPressTimer.current = id;
	}, [isTouchDevice, enabled, clearLongPress, longPressMs]);

	const onTouchEnd = useCallback(() => {
		clearLongPress();
	}, [clearLongPress]);

	const onKeyDown = useCallback(
		(e: KeyboardEvent<HTMLElement>) => {
			if (!enabled) return;

			if (e.key === "Enter") setOpen((open) => !open);
			if (e.key === "Escape") setOpen(false);
		},
		[enabled],
	);

	useEffect(
		() => () => {
			clearHoverOpen();
			clearLongPress();
		},
		[clearHoverOpen, clearLongPress],
	);

	useEffect(() => {
		if (!enabled) setOpen(false);
	}, [enabled]);

	return {
		open,
		handleOpenChange,
		triggerHandlers: { onMouseEnter, onMouseLeave, onTouchStart, onTouchEnd, onKeyDown },
	};
};

export { Popover, PopoverTrigger, PopoverContent };
