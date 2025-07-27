export function generateGradientColor(y: number, height: number): string {
  const hue = (y / height) * 60 + 200; // Blue to purple range
  return `hsla(${hue}, 70%, 60%, 0.8)`;
}
