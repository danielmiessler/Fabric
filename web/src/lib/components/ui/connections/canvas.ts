export function createParticleGradient(
    ctx: CanvasRenderingContext2D,
    x: number,
    y: number,
    size: number,
    color: string
  ): CanvasGradient {
    const gradient = ctx.createRadialGradient(x, y, 0, x, y, size);
    gradient.addColorStop(0, color);
    gradient.addColorStop(1, 'transparent');
    return gradient;
  }
