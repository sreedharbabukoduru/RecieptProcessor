# Build stage
FROM eclipse-temurin:17-jdk-alpine as builder
WORKDIR /build
COPY receiptprocessor.jar app.jar

# Runtime stage
FROM eclipse-temurin:17-jre-alpine
WORKDIR /app

# Add metadata
LABEL maintainer="Developer" \
      description="Receipt Processor Service" \
      version="1.0"

# Copy jar from builder stage
COPY --from=builder /build/app.jar ./app.jar

# Configure Java options
ENV JAVA_OPTS="-Xms512m -Xmx512m -XX:+UseG1GC"

# Expose application port
EXPOSE 3000

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget -q --spider http://localhost:3000/actuator/health || exit 1

# Run application
ENTRYPOINT ["sh", "-c", "java ${JAVA_OPTS} -jar app.jar"]
