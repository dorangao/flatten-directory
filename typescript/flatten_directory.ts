// flatten_directory.ts
import * as fs from 'fs';
import * as path from 'path';

const [,, sourceDir, targetDir] = process.argv;

if (!sourceDir || !targetDir) {
    console.error('Usage: node flatten_directory.js <source_directory> <target_directory>');
    process.exit(1);
}

const flattenDirectory = (sourceDir: string, targetDir: string) => {
    if (!fs.existsSync(sourceDir)) {
        console.error('Error: Source directory does not exist.');
        process.exit(1);
    }

    fs.mkdirSync(targetDir, { recursive: true });

    const flatten = (dir: string) => {
        const entries = fs.readdirSync(dir, { withFileTypes: true });
        entries.forEach((entry) => {
            const fullPath = path.join(dir, entry.name);
            if (entry.isFile()) {
                const targetPath = path.join(targetDir, entry.name);
                fs.copyFileSync(fullPath, targetPath);
            } else if (entry.isDirectory()) {
                flatten(fullPath);
            }
        });
    };

    flatten(sourceDir);
    console.log(`All files have been extracted to ${targetDir} without preserving folder structure.`);
};

flattenDirectory(sourceDir, targetDir);
