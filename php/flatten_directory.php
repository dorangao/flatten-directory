<?php
// flatten_directory.php

function flattenDirectory(string $sourceDir, string $targetDir): void {
    if (!is_dir($sourceDir)) {
        fwrite(STDERR, "Error: Source directory does not exist." . PHP_EOL);
        exit(1);
    }

    if (!is_dir($targetDir)) {
        mkdir($targetDir, 0777, true);
    }

    $files = new RecursiveIteratorIterator(
        new RecursiveDirectoryIterator($sourceDir, RecursiveDirectoryIterator::SKIP_DOTS),
        RecursiveIteratorIterator::SELF_FIRST
    );
    foreach ($files as $file) {
        if ($file->isFile()) {
            $targetPath = $targetDir . DIRECTORY_SEPARATOR . $file->getBasename();
            copy($file->getPathname(), $targetPath);
        }
    }
    echo "All files have been extracted to $targetDir without preserving folder structure." . PHP_EOL;
}

if ($argc !== 3) {
    fwrite(STDERR, "Usage: php flatten_directory.php <source_directory> <target_directory>" . PHP_EOL);
    exit(1);
}

flattenDirectory($argv[1], $argv[2]);
