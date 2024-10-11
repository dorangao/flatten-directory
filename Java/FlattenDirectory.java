// FlattenDirectory.java
import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardCopyOption;

public class FlattenDirectory {
    public static void flattenDirectory(String sourceDir, String targetDir) {
        File source = new File(sourceDir);
        File target = new File(targetDir);

        if (!source.exists() || !source.isDirectory()) {
            System.err.println("Error: Source directory does not exist or is not a directory.");
            System.exit(1);
        }

        if (!target.exists()) {
            target.mkdirs();
        }

        flatten(source, target);
        System.out.println("All files have been extracted to " + targetDir + " without preserving folder structure.");
    }

    private static void flatten(File sourceDir, File targetDir) {
        for (File entry : sourceDir.listFiles()) {
            if (entry.isFile()) {
                try {
                    Files.copy(entry.toPath(), Path.of(targetDir.getPath(), entry.getName()), StandardCopyOption.REPLACE_EXISTING);
                } catch (IOException e) {
                    System.err.println("Error copying file: " + entry.getPath());
                }
            } else if (entry.isDirectory()) {
                flatten(entry, targetDir);
            }
        }
    }

    public static void main(String[] args) {
        if (args.length != 2) {
            System.out.println("Usage: java FlattenDirectory <source_directory> <target_directory>");
            System.exit(1);
        }
        flattenDirectory(args[0], args[1]);
    }
}
