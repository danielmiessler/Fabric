import Root from "./button.svelte";

const buttonVariants = {
	base: "inline-flex items-center justify-center whitespace-nowrap rounded-full text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 disabled:pointer-events-none disabled:opacity-50",
	variants: {
		variant: {
			default: "bg-primary text-primary-foreground hover:bg-primary/90 rounded-full border shadow-sm",
			destructive: "bg-destructive text-destructive-foreground hover:bg-destructive/90 shadow-sm",
			outline: "border-input bg-background hover:bg-accent hover:text-accent-foreground border shadow-sm",
			secondary: "bg-secondary text-secondary-foreground hover:bg-secondary/80 rounded-full border variant-glass-secondary shadow-sm",
			ghost: "hover:bg-accent hover:text-accent-foreground",
			link: "text-primary underline-offset-4 hover:underline"
		},
		size: {
			default: "h-9 px-4 py-2",
			sm: "h-8 rounded-md px-3 text-xs",
			lg: "h-10 rounded-md px-8",
			icon: "h-9 w-9"
		}
	},
	defaultVariants: {
		variant: "default",
		size: "default"
	}
};

export {
	Root,
	Root as Button,
	buttonVariants
};
