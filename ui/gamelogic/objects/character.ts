import { Spritesheets } from "../enums";

export class Character extends Phaser.Physics.Arcade.Sprite {
    id: string | null = null;
    isDead: boolean = false;

    constructor(scene: Phaser.Scene, x: number, y: number, tint: number | null = null) {
        super(scene, x, y, Spritesheets.Main, 4);

        if (tint !== null)
            this.setTint(tint);

        scene.add.existing(this);
        scene.physics.add.existing(this);

        this.body?.setSize(7, 14, true);
        this.setCollideWorldBounds(true);

        this.anims.create({
            key: 'walk',
            frames: this.anims.generateFrameNumbers(Spritesheets.Main, { frames: [ 12, 13, 14, 15 ] }),
            frameRate: 8,
            repeat: -1
        });

        this.anims.create({
            key: 'idle',
            frames: this.anims.generateFrameNumbers(Spritesheets.Main, { frames: [ 8, 9 ] }),
            frameRate: 8,
            repeat: -1
        });

        this.anims.create({
            key: 'jump',
            frames: this.anims.generateFrameNumbers(Spritesheets.Main, { frames: [ 16, 17 ] }),
            frameRate: 4,
            repeat: -1
        });

        this.anims.create({
            key: 'die',
            frames: this.anims.generateFrameNumbers(Spritesheets.Main, { frames: [ 20, 21, 22, 23 ] }),
            frameRate: 2,
        });
    }

    update(time: number, delta: number): void {
        super.update(time, delta);
    }

    characterDie(): void {
        this.isDead = true;
        this.play('die');
    }

    spawn(x: number, y: number): void {
        this.isDead = false;
        this.play('idle');
        this.setPosition(x, y);
    }
}
