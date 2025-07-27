import type { Particle } from './particle';
import { generateGradientColor } from './colors';

export class ParticleSystem {
  private particles: Particle[] = [];
  private width: number;
  private height: number;
  private mouseX: number = 0;
  private mouseY: number = 0;
  private targetMouseX: number = 0;
  private targetMouseY: number = 0;

  constructor(
    private readonly count: number,
    private readonly baseSize: number,
    private readonly speed: number,
    width: number,
    height: number
  ) {
    this.width = width;
    this.height = height;
    this.mouseX = width / 2;
    this.mouseY = height / 2;
    this.targetMouseX = this.mouseX;
    this.targetMouseY = this.mouseY;
    this.initParticles();
  }

  private initParticles(): void {
    this.particles = [];
    for (let i = 0; i < this.count; i++) {
      // Distribute particles across the entire width
      const x = Math.random() * this.width;
      // Distribute particles vertically around the middle with some variation
      const yOffset = (Math.random() - 0.5) * 100;
      
      this.particles.push({
        x,
        y: this.height / 2 + yOffset,
        baseY: this.height / 2 + yOffset,
        speed: (Math.random() - 0.5) * this.speed * 0.5, // Reduced base speed
        angle: Math.random() * Math.PI * 2,
        size: this.baseSize * (0.8 + Math.random() * 0.4),
        color: generateGradientColor(this.height / 2 + yOffset, this.height),
        velocityX: (Math.random() - 0.5) * this.speed // Reduced initial velocity
      });
    }
  }

  public updateDimensions(width: number, height: number): void {
    this.width = width;
    this.height = height;
    this.mouseX = width / 2;
    this.mouseY = height / 2;
    this.targetMouseX = this.mouseX;
    this.targetMouseY = this.mouseY;
    this.initParticles();
  }

  public updateMousePosition(x: number, y: number): void {
    this.targetMouseX = x;
    this.targetMouseY = y;
  }

  public update(): void {
    // Smooth mouse movement
    this.mouseX += (this.targetMouseX - this.mouseX) * 0.05; // Slower mouse tracking
    this.mouseY += (this.targetMouseY - this.mouseY) * 0.05;

    this.particles.forEach(particle => {
      // Update horizontal position with constant motion
      particle.x += particle.velocityX;
      
      // Wave motion
      particle.angle += particle.speed;
      const waveAmplitude = 30 * (this.mouseY / this.height); // Reduced amplitude
      const frequencyFactor = (this.mouseX / this.width);
      
      // Calculate vertical position with wave effect
      particle.y = particle.baseY + 
        Math.sin(particle.angle * frequencyFactor + particle.x * 0.01) * // Slower wave
        waveAmplitude;
      
      // Update particle color based on position
      particle.color = generateGradientColor(particle.y, this.height);
      
      // Screen wrapping with position preservation
      if (particle.x < 0) {
        particle.x = this.width;
        particle.baseY = this.height / 2 + (Math.random() - 0.5) * 100;
      }
      if (particle.x > this.width) {
        particle.x = 0;
        particle.baseY = this.height / 2 + (Math.random() - 0.5) * 100;
      }

      // Very subtle velocity adjustment to maintain spread
      if (Math.abs(particle.velocityX) < 0.1) {
        particle.velocityX += (Math.random() - 0.5) * 0.02;
      }
      
      // Gentle velocity dampening
      particle.velocityX *= 0.99;
    });
  }

  public getParticles(): Particle[] {
    return this.particles;
  }
}
