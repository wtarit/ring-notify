interface EmptyStateProps {
  title: string;
  description: string;
  action?: React.ReactNode;
}

export function EmptyState({ title, description, action }: EmptyStateProps) {
  return (
    <div className="text-center py-16">
      <h3 className="text-lg font-semibold">{title}</h3>
      <p className="mt-2 opacity-70">{description}</p>
      {action && <div className="mt-4">{action}</div>}
    </div>
  );
}
