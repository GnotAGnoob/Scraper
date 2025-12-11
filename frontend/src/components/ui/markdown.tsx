import ReactMarkdown from "react-markdown";
import { cn } from "@/lib/utils";

interface MarkdownProps {
	children: string;
	className?: string;
}

export const Markdown = ({ children, className }: MarkdownProps) => {
	return (
		<div
			className={cn(
				"prose prose-neutral prose-sm dark:prose-invert max-w-none",
				"prose-headings:font-black",
				"prose-h1:mt-2 prose-h1:mb-1 prose-h1:text-2xl",
				"prose-h2:mt-2 prose-h2:mb-1 prose-h2:text-xl",
				"prose-h3:mt-1.5 prose-h3:mb-0.5 prose-h3:text-lg",
				"prose-h4:mt-1.5 prose-h4:mb-0.5 prose-h4:text-base",
				"prose-p:mt-1 prose-p:mb-1 prose-p:leading-snug",
				"prose-ul:mt-1 prose-ul:mb-1",
				"prose-ol:mt-1 prose-ol:mb-1",
				"prose-li:my-0",
				"prose-blockquote:mt-1 prose-blockquote:mb-1",
				className,
			)}
		>
			<ReactMarkdown>{children}</ReactMarkdown>
		</div>
	);
};
